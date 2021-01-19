package model

type GetRoleGroupRequest struct {
	GroupId int `json:"group_id"`
}

type GetRoleGroupResponse struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
	RoleGroup struct {
		GroupName string `json:"group_name"`
		Roles     []struct {
			RoleId   int    `json:"role_id"`
			RoleName string `json:"role_name"`
		} `json:"roles"`
	} `json:"role_group"`
}

type GetRoleUserRequest struct {
	RoleId int `json:"role_id"`
}

type GetRoleUserResponse struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
	Result    struct {
		HasMore    bool `json:"hasMore"`
		NextCursor int  `json:"nextCursor"`
		User       []struct {
			UserId string `json:"userid"`
			Name   string `json:"name"`
		} `json:"list"`
	} `json:"result"`
}

type RoleListResponse struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
	Result    struct {
		HasMore bool `json:"has_more"`
		List    []struct {
			Name    string `json:"name"`
			GroupId int    `json:"groupId"`
			Roles   []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"roles"`
		} `json:"list"`
	} `json:"result"`
}
