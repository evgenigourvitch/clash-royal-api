package crawler

import (
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/evgenigourvitch/clash-royal-api/objects"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	cNilResponseErr          = errors.New("nil response")
	cUndefinedResponseObject = errors.New("got undefined object in response")
	cMaxNumOfRetries         = 10
)

type responseChannelObject struct {
	obj    interface{}
	result interface{}
	err    error
}

type crawler struct {
	httpClient *http.Client
	reqHeaders http.Header
	swaggerUrl string
}

func (c *crawler) GetSwaggerUrl() string {
	return c.swaggerUrl
}

func (c *crawler) RequestAsync(url string, responseObjectType objects.ResponseType, resultsChannel chan responseChannelObject, id interface{}) {
	res, err := c.Request(url, responseObjectType, 0)
	if err != nil {
		resultsChannel <- responseChannelObject{id, nil, err}
		return
	}
	resultsChannel <- responseChannelObject{id, res, nil}
}

func (c *crawler) Request(url string, responseObjectType objects.ResponseType, retryCnt int) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.reqHeaders
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusTooManyRequests {
		if retryCnt <= cMaxNumOfRetries {
			return c.Request(url, responseObjectType, retryCnt+1)
		}
		return nil, fmt.Errorf("failed to fetch url %s after %d retries\n", url, retryCnt)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code: %d, url: %s", res.StatusCode, url)
	}
	bytesArr, err := c.readBody(res)
	if err != nil {
		return nil, err
	}
	return c.parseResponse(bytesArr, responseObjectType)
}

func (c *crawler) GetBattles(clanMembers *objects.PlayersResponse) {
	requestResultChan := make(chan responseChannelObject, len(clanMembers.Players))
	for _, member := range clanMembers.Players {
		tag := strings.TrimLeft(member.Tag, "#")
		go c.RequestAsync(c.swaggerUrl+"players/%23"+tag+"/battlelog", objects.EResponseTypeBattles, requestResultChan, member)
	}
	for i := 0; i < cap(requestResultChan); i++ {
		select {
		case requestResult := <-requestResultChan:
			if requestResult.err != nil {
				member := requestResult.obj.(*objects.Player)
				fmt.Printf("failed to get battles for player: %s, error: %+v\n", member.Name, requestResult.err)
			} else {
				battles := requestResult.result.([]*objects.Battle)
				fmt.Printf("battles: %+v\n", battles)
			}
		}
	}
}

func (c *crawler) parseResponse(bytesArr []byte, responseObjectType objects.ResponseType) (interface{}, error) {
	switch responseObjectType {
	case objects.EResponseTypePlayersList:
		return objects.ParsePlayers(bytesArr)
	case objects.EResponseTypeBattles:
		return objects.ParseBattles(bytesArr)
	default:
		return nil, cUndefinedResponseObject
	}
	return nil, nil
}

func (c *crawler) readBody(res *http.Response) ([]byte, error) {
	if res == nil {
		return nil, cNilResponseErr
	}
	body := res.Body
	defer body.Close()
	var reader io.ReadCloser
	var err error
	if res.Header.Get("Content-Encoding") == "gzip" {
		if reader, err = gzip.NewReader(body); err != nil {
			return nil, err
		}
		defer reader.Close()

	} else {
		reader = body
	}
	bytesArr, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return bytesArr, nil
}

func NewCralwer(swaggerUrl, authToken string) (*crawler, error) {
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{MaxIdleConnsPerHost: 256, MaxIdleConns: 256, IdleConnTimeout: time.Minute * 2}
	headers := http.Header{"Accept": []string{"application/json"},
		"Authorization":   []string{"Bearer " + authToken},
		"Accept-Encoding": []string{"gzip"},
	}
	c := &crawler{httpClient, headers, swaggerUrl}
	return c, nil
}

func initToken() (string, error) {
	b, err := ioutil.ReadFile("./.token")
	if err != nil {
		fmt.Printf("Failed to get auth token, error: %+v", err)
		return "", err
	}
	return strings.Trim(string(b), "\n"), nil
}
