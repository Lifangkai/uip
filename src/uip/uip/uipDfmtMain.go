package main

import (
	"net/http"
	"fmt"

	"uip/attached"
	"uip/common"
)

/**
*	添加数据方法
*/
func mainInsertValueMethod(w http.ResponseWriter, data attached.UipDfmtMain) {
	frame.Log.Write("***mainInsertValueMethod***")

	//校验数据
	if ok := attached.MainFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

  if data.FmtCode == "" {
  		//获取SN
		var errs error
		data.FmtCode, errs = getSnsNew(data.GroupId ,"fmtCode")
		if errs != nil {
			fmt.Println(errs.Error())
			response.Code = common.ErrorDiscoverCheckId
			response.Msg = common.ErrorDiscoverCheckMsg
			return
		}
 	 } 

	//添加数据
	if err := mainAddMehod(data); err != nil {
		fmt.Println(err.Error())
		if err.Error()=="数据已存在"{
			response.Code = common.ErrorDataExistsErrId
			response.Msg = common.ErrorDataExistsErrMsg
			return
		}
		response.Code = common.ErrorAddInfoId
		response.Msg = common.ErrorAddInfoMsg
		return
	}

	//成功响应
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKInsertMsg
}

/**
*	修改数据方法
*/
func mainUpdateValueMethod(w http.ResponseWriter, data attached.UipDfmtMain) {
	frame.Log.Write("***mainUpdateValueMethod***")

	//校验数据
	if ok := attached.MainFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//修改
	if err := mainUpdateMethod(data); err != nil {
		fmt.Println(err.Error())
		if err.Error()=="数据不存在"{
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorUpdataInfoFailedId
		response.Msg = common.ErrorUpdataInfoFailedMsg
		return
	}

	//成功响应
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKModifyMsg
}

/**
*	删除数据方法
*/
func mainDeleteValueMethod(w http.ResponseWriter, data attached.UipDfmtMain) {
	frame.Log.Write("***mainDeleteValueMethod***")

	//数据过滤
	if data.GroupId == "" || data.FmtCode == "" || len(data.GroupId) != 8 || len(data.FmtCode) > 32 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//删除数据
	if err := mainDelMethod(data); err != nil {
		fmt.Println(err)
		if err.Error()=="数据不存在"{
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorDelErrId
		response.Msg = common.ErrorDelErrMsg
		return
	}

	//成功
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKDeleteMsg
}

/**
*	查询主表与子表的数据方法
*/
func mainQueryValueMethod(w http.ResponseWriter, r *http.Request) {
	frame.Log.Write("***mainQueryValueMethod***")

	//接受数据并校验
	if r.Form["groupId"][0] == "" || len(r.Form["groupId"]) != 1 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	if r.Form["fmtCode"][0] == "" || len(r.Form["fmtCode"]) != 1 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//查询的条件
	key := r.Form["groupId"][0] + r.Form["fmtCode"][0]

	//查询主表
	resultUipDfmtMain, err := mainQueryOneMethod(key)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error()=="数据不存在"{
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorQueryErrId
		response.Msg = common.ErrorQueryErrMsg
		return
	}

	//查询子表
	resultUipDfmtSub, err1 := querySubAllCondtion(key)
	if err1 != nil {
		fmt.Println(err1.Error())
		if err.Error()=="数据为空"{
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorQueryErrId
		response.Msg = common.ErrorQueryErrMsg
		return
	}

	//返回的数据结构
	type mainAndSub struct {
		FmtName string                `json:"fmtName"`
		FmtType string                `json:"fmtType"`
		ExtInfo string                `json:"extInfo"`
		Cmt     string                `json:"cmt"`
		Fields  []attached.UipDfmtSub `json:"fields"`
	}

	//赋值
	finalValue := mainAndSub{
		resultUipDfmtMain.FmtName,
		resultUipDfmtMain.FmtType,
		resultUipDfmtMain.ExtInfo,
		resultUipDfmtMain.Cmt,
		resultUipDfmtSub,
	}

	//返回
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKSelectMsg
	response.Data = finalValue
}
