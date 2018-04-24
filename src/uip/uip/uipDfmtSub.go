package main

import (
	"net/http"
	"uip/attached"
	"fmt"
	"uip/common"
)

/*
	添加数据方法
*/
func subInsertValueMethod(w http.ResponseWriter,r *http.Request,data attached.UipDfmtSub)  {
	fmt.Println("***mainInsertValueMethod***")

	//响应状态
	defer response.Answer(w)

	//校验数据
	if ok:=attached.SubFilter(data,w);!ok{
		response.Code=common.CheckErrorId
		response.Code=common.CheckErrorMsg
		return
	}

}

/*
	修改数据方法
*/
func subUpdateValueMethod(w http.ResponseWriter,r *http.Request,data attached.UipDfmtSub)  {
	fmt.Println("***mainUpdateValueMethod***")

	//响应状态
	defer response.Answer(w)

	//校验数据
	if ok:=attached.SubFilter(data,w);!ok{
		response.Code=common.CheckErrorId
		response.Code=common.CheckErrorMsg
		return
	}
}

/*
	删除数据方法
*/
func subDeleteValueMethod(w http.ResponseWriter,r *http.Request,data attached.UipDfmtSub)  {
	fmt.Println("***mainDeleteValueMethod***")
}

