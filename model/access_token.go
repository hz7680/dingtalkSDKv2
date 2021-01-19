package model

import (
	"time"
)

type AccessTokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type AccessToken struct {
	accessToken string
	expiredAt   time.Time
}

func (a *AccessToken) IsExpired() bool {
	return a.accessToken == "" || a.expiredAt.Unix()-100 > time.Now().Unix()
}
func (a *AccessToken) SetAccessToken(token string, expiresIn int) {
	a.accessToken = token
	a.expiredAt = time.Now().Add(time.Duration(expiresIn) * time.Second)
}

func (a *AccessToken) GetAccessToken() string {
	return a.accessToken
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		accessToken: "",
		expiredAt:   time.Now().Add(-1 * time.Second),
	}
}
