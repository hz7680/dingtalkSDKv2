package client

import (
	"dingtalkSdkV2/model"
	"dingtalkSdkV2/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) GetRoleGroup(groupId int) (resp *model.GetRoleGroupResponse, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/role/getrolegroup?access_token=%s", token)
	body, err := utils.HttpPostJson(url, &model.GetRoleGroupRequest{GroupId: groupId})
	if err != nil {
		return
	}
	fmt.Println(string(body))
	resp = &model.GetRoleGroupResponse{}
	err = json.Unmarshal(body, resp)
	return
}
