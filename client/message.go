package client

import (
	"dingtalkSdkV2/model"
	"dingtalkSdkV2/utils"
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) SendCorpConversionToUsers(userIds []string, msg interface{}) (resp *model.SendCorpConversationResponse, err error) {
	request := model.SendCorpConversationByUserIdRequest{
		AgentId:    c.agentId,
		UseridList: strings.Join(userIds, ","),
		Msg:        msg,
	}
	return c.sendCorpConversion(request)
}

func (c *Client) SendCorpConversionToDepts(deptIds []string, msg interface{}) (resp *model.SendCorpConversationResponse, err error) {
	request := model.SendCorpConversationByDeptIdListRequest{
		AgentId:    c.agentId,
		DeptIdList: strings.Join(deptIds, ","),
		Msg:        msg,
	}
	return c.sendCorpConversion(request)
}

func (c *Client) SendCorpConversionToAllUsers(msg interface{}) (resp *model.SendCorpConversationResponse, err error) {
	request := model.SendCorpConversationToAllUserRequest{
		AgentId:   c.agentId,
		ToAllUser: true,
		Msg:       msg,
	}
	return c.sendCorpConversion(request)
}

func (c *Client) sendCorpConversion(req interface{}) (resp *model.SendCorpConversationResponse, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", token)
	body, err := utils.HttpPostJson(url, req)
	if err != nil {
		return
	}
	resp = &model.SendCorpConversationResponse{}
	err = json.Unmarshal(body, resp)
	return

}
