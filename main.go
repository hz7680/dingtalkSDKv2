package main

import (
	"dingtalkSdkV2/client"
	"dingtalkSdkV2/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		all, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println("got an error when read body : ", err)
			return
		}
		temp := string(all)
		fmt.Println(temp)
		var req model.CallBackRequest
		err = json.Unmarshal(all, &req)
		if err != nil {
			fmt.Println("got an error when json unmarshal : ", err)
			return
		}
		req.Nonce = request.URL.Query().Get("nonce")
		req.Signature = request.URL.Query().Get("signature")
		req.TimeStamp = request.URL.Query().Get("timestamp")
		decryptMsg, err := dingtalkClient.Decrypt(&req)
		if err != nil {
			fmt.Println("got an error when decrypt : ", err)
		}
		fmt.Println(decryptMsg)
		//if decryptMsg["EventType"] == "check_url"{
		resp, err := dingtalkClient.GenerateSuccessResp(&req)
		if err != nil {
			fmt.Println("got an error when generate success response :", err)
		}
		marshalBytes, err := json.Marshal(resp)
		writer.Write(marshalBytes)
		//}
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	err := http.ListenAndServe(":8001", nil)
	fmt.Println(err)

}
