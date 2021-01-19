package model

type SendCorpConversationByUserIdRequest struct {
	AgentId    int         `json:"agent_id"`
	UseridList string      `json:"userid_list"`
	Msg        interface{} `json:"msg"`
}

type SendCorpConversationByDeptIdListRequest struct {
	AgentId    int         `json:"agent_id"`
	DeptIdList string      `json:"dept_id_list"`
	Msg        interface{} `json:"msg"`
}

type SendCorpConversationToAllUserRequest struct {
	AgentId   int         `json:"agent_id"`
	ToAllUser bool        `json:"to_all_user"`
	Msg       interface{} `json:"msg"`
}

type MsgSingleActionCard struct {
	Msgtype    string `json:"msgtype"`
	ActionCard struct {
		Title       string `json:"title"`
		Markdown    string `json:"markdown"`
		SingleTitle string `json:"single_title"`
		SingleUrl   string `json:"single_url"`
	} `json:"action_card"`
}

func NewMsgSingleActionCard(title, markdown, singleTitle, singleUrl string) *MsgSingleActionCard {
	return &MsgSingleActionCard{
		Msgtype: "action_card",
		ActionCard: struct {
			Title       string `json:"title"`
			Markdown    string `json:"markdown"`
			SingleTitle string `json:"single_title"`
			SingleUrl   string `json:"single_url"`
		}{
			Title:       title,
			Markdown:    markdown,
			SingleTitle: singleTitle,
			SingleUrl:   singleUrl,
		},
	}
}

type SendCorpConversationResponse struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
	TaskId    int    `json:"task_id"`
}
