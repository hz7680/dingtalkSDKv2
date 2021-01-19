package client

import (
	"dingtalkSdkV2/model"
	"dingtalkSdkV2/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) GetUserDetail(userId string) (resp *model.UserDetailResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/get?access_token=%s", token)
	body, err := utils.HttpPostJson(url, &model.UserDetailReq{UserId: userId})
	if err != nil {
		return
	}
	//fmt.Println(string(body))
	resp = &model.UserDetailResp{}
	err = json.Unmarshal(body, resp)
	return
}
