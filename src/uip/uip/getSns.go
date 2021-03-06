package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"errors"
	
	"uip/common"
)

//调用sns服务获取序列号
func getSnsNew(groupId string, snsType string) (string, error) {
	common.Frame.Log.Write("in getSnsNew(): groupid = " + groupId + "snstype = " + snsType)
	
	//请求sns的路径(网址)
	url := "http://"+common.AExt.SNSIP + ":" + common.AExt.SNSPORT + "/snList?com=LIST&groupId="+ groupId +"&snId=" + snsType + "&operatorId=12000&bs=1"
	common.Frame.Log.Write("url = " + url)
	
	//进行http请求 获取请求数据
	resp, err := http.Get(url)
	if err != nil {
		//错误处理
		return "", err
	}
	//读取响应主体内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//错误处理
		return "", err
	}
	//定义响应的结构体用于解析收到的响应数据
	type sn struct {
		Sn string `json:"Sn"`
	}
	
	var snsInfo struct {
		Code string `json:"Code"`
		Msg  string `json:"Msg"`
		Data []sn   `json:"Data"`
	}
	
	//解析响应数据
	err = json.Unmarshal(respBody, &snsInfo)
	if err != nil {
		common.Frame.Log.Write("json.Unmarshal error: " + err.Error())
		//错误处理
		return "", err
	}
	
	//判断响应code
	if snsInfo.Code != "200000" {
		//不是200000的话 返回错误信息
		return "", errors.New(snsInfo.Msg)
	}
	
	//去空格处理
	snString := strings.Replace(snsInfo.Data[0].Sn, " ", "", -1)
	
	return snString, nil
}
