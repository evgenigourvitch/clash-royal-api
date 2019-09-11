package login

import (
	"bytes"
	"fmt"
	"github.com/evgenigourvitch/clash-royal-api/objects"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	cPostDataTemplate = `{"email":"%s","password":"%s"}`
	cLoginURL         = "https://developer.clashroyale.com/api/login"
)

type loginService struct {
	httpClient *http.Client
	reqHeaders http.Header
	postData   []byte
	loginURL   *url.URL
}

func (ls *loginService) Login() (*objects.LoginResponse, error) {
	req := &http.Request{URL: ls.loginURL,
		Method:        http.MethodPost,
		Body:          ioutil.NopCloser(bytes.NewReader(ls.postData)),
		ContentLength: int64(len(ls.postData)),
		Header:        ls.reqHeaders,
	}
	res, err := ls.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code: %d, on login request", res.StatusCode)
	}
	return ls.parseLoginResponse(res)
}

func (ls *loginService) parseLoginResponse(res *http.Response) (*objects.LoginResponse, error) {
	body := res.Body
	defer body.Close()
	bytesArr, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return objects.ParseLoginResponse(bytesArr)
}

func NewLoginService(credentials *objects.Credentials) *loginService {
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{MaxIdleConnsPerHost: 256, MaxIdleConns: 256, IdleConnTimeout: time.Minute * 2}
	headers := http.Header{"Content-Type": []string{"application/json"}}
	postData := []byte(fmt.Sprintf(cPostDataTemplate, credentials.Email, credentials.Password))
	parsedUrl, err := url.Parse(cLoginURL)
	if err != nil {
		panic(err)
	}
	return &loginService{httpClient, headers, postData, parsedUrl}
}
