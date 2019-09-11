package objects

import (
	"encoding/json"
	"io/ioutil"
)

type Credentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewCredentials() (*Credentials, error) {
	b, err := ioutil.ReadFile("./.credentials.json")
	if err != nil {
		return nil, err
	}
	result := new(Credentials)
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil

}
