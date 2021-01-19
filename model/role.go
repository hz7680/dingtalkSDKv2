package model

type GetRoleGroupRequest struct {
	GroupId int `json:"group_id"`
}

type GetRoleGroupResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	RequestId string `json:"request_id"`
	RoleGroup struct{
		GroupName string `json:"group_name"`
		Roles []struct{
			RoleId int `json:"role_id"`
			RoleName string `json:"role_name"`
		} `json:"roles"`
	} `json:"role_group"`
}
