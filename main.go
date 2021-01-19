package main

import (
	"dingtalkSdkV2/client"
	"fmt"
)

// access_token: 9151da54523233968275463580ccfaff
// group_id: 1814596691
func main() {
	corpId := "ding55eb045fd2ec4e54f5bf40eda33b7ba0"
	agentId := "1068530230"
	appKey := "dingz05ymbymlkurako3"
	appSecret := "LbLg-hEvFdGFFNPBiSfpUKt_v2kZybSHCZPfdLLSVy_gs0M5iHQsWKMagxisBCd-"
	aesKey := "Fqnxq5sawertHvTwjAi4TjUnsAj9gAlsdNemXjfSNDU"
	token := "vl1xnJrdsXF06QYWs77"
	dingtalkClient := client.NewClient(corpId, agentId, appKey, appSecret, aesKey, appKey, token)
	//accessToken, err := dingtalkClient.GetAccessToken()
	//fmt.Println(accessToken)
	//fmt.Print(err)
	err := dingtalkClient.GetRoleGroup(1814596691)
	fmt.Println(err)
}
