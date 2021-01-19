package client

import (
	"dingtalkSdkV2/crypto"
	"dingtalkSdkV2/model"
)

type Client struct {
	corpId      string
	agentId     int
	appKey      string
	appSecret   string
	aesKey      string
	key         string
	accessToken *model.AccessToken
	cryptor     *crypto.Crypto
}

func NewClient(corpId string, agentId int, appKey, appSecret, aesKey, key, token string) *Client {
	return &Client{
		corpId:      corpId,
		agentId:     agentId,
		appKey:      appKey,
		appSecret:   appSecret,
		aesKey:      aesKey,
		key:         key,
		accessToken: model.NewAccessToken(),
		cryptor:     crypto.NewCrypto(token, aesKey, key),
	}
}
