package client

import (
	"dingtalkSdkV2/model"
	"dingtalkSdkV2/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) GetUserIdByCode(code string) (resp *model.GetUserIdByCodeResponse, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://oapi.dingtalk.com/user/getuserinfo?access_token=%s&code=%s", token, code)
	body, err := utils.HttpGet(url)

	resp = &model.GetUserIdByCodeResponse{}
	err = json.Unmarshal(body, resp)
	return
}
