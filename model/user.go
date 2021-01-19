package model

type UserDetailReq struct {
	UserId string `json:"userid"`
}
type UserDetailResp struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
	Result    struct {
		Boss     bool   `json:"boss"`
		Unionid  string `json:"unionid"`
		RoleList []struct {
			GroupName string `json:"group_name"`
			Name      string `json:"name"`
			Id        int    `json:"id"`
		} `json:"role_list"`
		Mobile        string `json:"mobile"`
		Active        bool   `json:"active"`
		Admin         bool   `json:"admin"`
		Remark        string `json:"remark"`
		Telephone     string `json:"telephone"`
		Avatar        string `json:"avatar"`
		HideMobile    bool   `json:"hide_mobile"`
		Title         string `json:"title"`
		UserId        string `json:"userid"`
		Senior        bool   `json:"senior"`
		WorkPlace     string `json:"work_place"`
		DeptOrderList []struct {
			DeptId int `json:"dept_id"`
			Order  int `json:"order"`
		} `json:"dept_order_list"`
		RealAuthed   bool   `json:"real_authed"`
		Name         string `json:"name"`
		DeptIdList   []int  `json:"dept_id_list"`
		JobNumber    string `json:"job_number"`
		StateCode    string `json:"state_code"`
		Email        string `json:"email"`
		LeaderInDept []struct {
			Leader bool `json:"leader"`
			DeptId int  `json:"dept_id"`
		} `json:"leader_in_dept"`
	} `json:"result"`
}
