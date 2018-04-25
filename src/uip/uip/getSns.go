package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"errors"
)

//调用sns服务获取序列号
func getSnsNew(snsType string) (string, error) {
	//请求sns的路径(网址)
	url := "http://172.16.0.14:31300/snList?com=LIST&groupId=12345678&snId=" + snsType + "&operatorId=12000&bs=1"
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
