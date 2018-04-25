package main

import (
	"net/http"
	"fmt"

	"uip/attached"
	"uip/common"
)

/*
	添加数据方法
*/
func mainInsertValueMethod(w http.ResponseWriter, data attached.UipDfmtMain) {
	fmt.Println("***mainInsertValueMethod***")

	//响应状态
	defer response.Answer(w)

	//校验数据
	if ok := attached.MainFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Code = common.CheckErrorMsg
		return
	}

	//获取SN
	//data.FmtCode,err:=getSn("")

	if err := mainAddMehod(data); err != nil {
		fmt.Println(err.Error())
		response.Code = common.ErrorAddInfoId
		response.Msg = common.ErrorAddInfoMsg
		return
	}

	//成功响应
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKInsertMsg
}

/*
	修改数据方法
*/
func mainUpdateValueMethod(w http.ResponseWriter, data attached.UipDfmtMain) {
	fmt.Println("***mainUpdateValueMethod***")

	//响应状态
	defer response.Answer(w)

	//校验数据
	if ok := attached.MainFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Code = common.CheckErrorMsg
		return
	}

	//修改
	if err := mainUpdateMehod(data); err != nil {
		response.Code = common.ErrorUpdataInfoFailedId
		response.Msg = common.ErrorUpdataInfoFailedMsg
		return
	}

	//成功响应
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKModifyMsg
}

/*
	删除数据方法
*/
func mainDeleteValueMethod(w http.ResponseWriter, data attached.UipDfmtMain) {
	fmt.Println("***mainDeleteValueMethod***")

	//响应状态
	defer response.Answer(w)

	//数据过滤
	if data.GroupId == "" || data.FmtCode == "" || len(data.GroupId) != 8 || len(data.FmtCode) > 32 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//删除数据
	if err := mainDelMethod(data); err != nil {
		response.Code = common.ErrorDelErrId
		response.Msg = common.ErrorDelErrMsg
		return
	}

	//成功
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKDeleteMsg
}

/*
	查询主表与子表的数据方法
*/
func mainQueryValueMethod(w http.ResponseWriter) {
	fmt.Println("***mainQueryValueMethod***")

}
