package main

import (
	"dingtalkSdkV2/client"
	"net/http"
)

var dingtalkClient *client.Client

// access_token: 9151da54523233968275463580ccfaff
// group_id: 1814596691
func main() {
	corpId := "ding55eb045fd2ec4e54f5bf40eda33b7ba0"
	agentId := 1068530230
	appKey := "dingz05ymbymlkurako3"
	appSecret := "LbLg-hEvFdGFFNPBiSfpUKt_v2kZybSHCZPfdLLSVy_gs0M5iHQsWKMagxisBCd-"
	aesKey := "Fqnxq5sawertHvTwjAi4TjUnsAj9gAlsdNemXjfSNDU"
	token := "vl1xnJrdsXF06QYWs77"
	dingtalkClient = client.NewClient(corpId, agentId, appKey, appSecret, aesKey, appKey, token)
	//accessToken, err := dingtalkClient.GetAccessToken()
	//fmt.Println(accessToken)
	//fmt.Print(err)
	//group, err := dingtalkClient.GetRoleGroup(1814596691)
	//fmt.Println(group)
	//fmt.Println(err)
	//users, err := dingtalkClient.SendCorpConversionToUsers([]string{"1919581426687495"}, model.NewMsgSingleActionCard("test", "### 您有新的待处理事项 \n * "+time.Now().Format("2006-01-02 15:04:05"), "立即查看", "http://www.baidu.com"))
	//fmt.Println(users)
	//fmt.Println(err)

	http.HandleFunc("/callback", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.ListenAndServe("8001",nil)

}
