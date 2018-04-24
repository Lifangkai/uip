package main

import (
	"net/http"
	"fmt"

	"uip/attached"
	"uip/common"
)

//type UipDfmtMain struct {
//	Cmt      string `json:"cmt"`
//	OperCode string `json:"operCode"`
//	Ctime    string `json:"ctime"`
//	ExtInfo  string `json:"extInfo"`
//	FmtType  string `json:"fmtType"`
//	FmtCode  string `json:"fmtCode"`
//	FmtName  string `json:"fmtName"`
//	GroupId  string `json:"groupId"`
//}
//
//type mainRequestJson struct {
//	Com  string     `json:"com"`
//	Data uipDfmtMain `json:"data"`
//}
//
//type mainResponseJson struct {
//	Code string `json:"code"`
//	Msg  string `json:"msg"`
//}
//type mainResponseJsonOne struct {
//	Code string     `json:"code"`
//	Msg  string     `json:"msg"`
//	Data uipDfmtMain `json:"data"`
//}
//type mainResponseJsonArray struct {
//	Code string       `json:"code"`
//	Msg  string       `json:"msg"`
//	Data []uipDfmtMain `json:"data"`
//}
//
//type requestStrings struct {
//	Com  string   `json:"com"`
//	Data []string `json:"data"`
//}

/*
	添加数据方法
*/
func mainInsertValueMethod(w http.ResponseWriter,r *http.Request,data attached.UipDfmtMain)  {
	fmt.Println("***mainInsertValueMethod***")

	//响应状态
	defer response.Answer(w)

	//校验数据
	if ok:=attached.MainFilter(data,w);!ok{
		response.Code=common.CheckErrorId
		response.Code=common.CheckErrorMsg
		return
	}

}

/*
	修改数据方法
*/
func mainUpdateValueMethod(w http.ResponseWriter,r *http.Request,data attached.UipDfmtMain)  {
	fmt.Println("***mainUpdateValueMethod***")

	//响应状态
	defer response.Answer(w)

	//校验数据
	if ok:=attached.MainFilter(data,w);!ok{
		response.Code=common.CheckErrorId
		response.Code=common.CheckErrorMsg
		return
	}
}

/*
	删除数据方法
*/
func mainDeleteValueMethod(w http.ResponseWriter,r *http.Request,data attached.UipDfmtMain)  {
	fmt.Println("***mainDeleteValueMethod***")
}

/*
	查询主表与子表的数据方法
*/
func mainQueryValueMethod(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("***mainQueryValueMethod***")
}