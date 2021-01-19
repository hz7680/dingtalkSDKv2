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
	//fmt.Println(string(body))
	resp = &model.GetRoleGroupResponse{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) GetRoleGroupList() (resp *model.RoleListResponse, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/role/list?access_token=%s", token)
	body, err := utils.HttpGet(url)
	if err != nil {
		return
	}
	//fmt.Println(string(body))
	resp = &model.RoleListResponse{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) GetRoleUser(roleId int) (resp *model.GetRoleUserResponse, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/role/simplelist?access_token=%s", token)
	body, err := utils.HttpPostJson(url, &model.GetRoleUserRequest{RoleId: roleId})
	if err != nil {
		return
	}
	//fmt.Println(string(body))
	resp = &model.GetRoleUserResponse{}
	err = json.Unmarshal(body, resp)
	return
}
