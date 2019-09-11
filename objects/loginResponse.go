package objects

import (
	"encoding/json"
	"fmt"
)

type status struct {
	Code    uint32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

type auth struct {
	UID   string `json:"uid,omitempty"`
	Token string `json:"token,omitempty"`
	UA    string `json:"ua,omitempty"`
	IP    string `json:"ip,omitempty"`
}

type developer struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Game          string `json:"game,omitempty"`
	Email         string `json:"email,omitempty"`
	Tier          string `json:"tier,omitempty"`
	AllowedScopes string `json:"allowedScopes,omitempty"`
	MaxCidrs      string `json:"maxCidrs,omitempty"`
	PrevLoginTs   string `json:"prevLoginTs,omitempty"`
	PrevLoginIp   string `json:"prevLoginIp,omitempty"`
	PrevLoginUa   string `json:"prevLoginUa,omitempty"`
}

type LoginResponse struct {
	TemporaryAPIToken       string     `json:"temporaryAPIToken,omitempty"`
	SessionExpiresInSeconds uint64     `json:"sessionExpiresInSeconds,omitempty"`
	SwaggerUrl              string     `json:"swaggerUrl,omitempty"`
	Status                  *status    `status:"swaggerUrl,omitempty"`
	Auth                    *auth      `json:"auth,omitempty"`
	Developer               *developer `json:"developer,omitempty"`
}

func (lr *LoginResponse) Print() {
	if lr == nil {
		return
	}
	if lr.Status != nil {
		fmt.Printf("%+v\n", *lr.Status)
	}
	if lr.Auth != nil {
		fmt.Printf("%+v\n", *lr.Auth)
	}
	if lr.Developer != nil {
		fmt.Printf("%+v\n", lr.Developer)
	}
	fmt.Printf("SwaggerUrl: %s, SessionExpiresInSeconds: %d, TemporaryAPIToken: %s\n", lr.SwaggerUrl, lr.SessionExpiresInSeconds, lr.TemporaryAPIToken)
}

func ParseLoginResponse(bytesArr []byte) (*LoginResponse, error) {
	result := new(LoginResponse)
	if err := json.Unmarshal(bytesArr, result); err != nil {
		return nil, err
	}
	return result, nil
}
