package model

type GetUserIdByCodeResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Userid   string `json:"userid"`
	Name     string `json:"name"`
	DeviceId string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
}
