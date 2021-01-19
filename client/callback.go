package client

import (
	"dingtalkSdkV2/model"
	"encoding/json"
)

func (c *Client) Decrypt(callBack *model.CallBackRequest) (decryptMsg map[string]interface{}, err error) {
	decryptMsgJson, err := c.cryptor.DecryptMsg(callBack.Signature, callBack.TimeStamp, callBack.Nonce, callBack.Encrypt)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(decryptMsgJson), &decryptMsg)
	if err != nil {
		return
	}
	return
}

func (c *Client) Encrypt(reply string, back *model.CallBackRequest) (encrypt string, signature string, err error) {
	return c.cryptor.EncryptMsg(reply, back.TimeStamp, back.Nonce)
}

func (c *Client) GenerateSuccessResp(callback *model.CallBackRequest) (resp *model.CallBackResponse, err error) {
	encrypt, signature, err := c.Encrypt("success", callback)
	if err != nil {
		return
	}
	resp = &model.CallBackResponse{
		MsgSignature: signature,
		TimeStamp:    callback.TimeStamp,
		Nonce:        callback.Nonce,
		Encrypt:      encrypt,
	}
	return
}