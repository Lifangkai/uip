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
func subInsertValueMethod(w http.ResponseWriter, data attached.UipDfmtSub) {
	fmt.Println("***mainInsertValueMethod***")

	//校验数据
	if ok := attached.SubFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//获取SN
	var errs error
	data.DtlCode, errs = getSnsNew(data.GroupId ,"fmtCode")
	if errs != nil {
		fmt.Println(errs.Error())
		response.Code = common.ErrorDiscoverCheckId
		response.Msg = common.ErrorDiscoverCheckMsg
		return
	}

	//添加数据
	if err := subAddMehod(data); err != nil {
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

/*
	修改数据方法
*/
func subUpdateValueMethod(w http.ResponseWriter, data attached.UipDfmtSub) {
	fmt.Println("***mainUpdateValueMethod***")

	//校验数据
	if ok := attached.SubFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//修改
	if err := subUpdateMehod(data); err != nil {
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

/*
	删除数据方法
*/
func subDeleteValueMethod(w http.ResponseWriter, data attached.UipDfmtSub) {
	fmt.Println("***mainDeleteValueMethod***")

	//数据过滤
	if data.GroupId == "" || data.FmtCode == "" || data.DtlCode == "" || len(data.GroupId) != 8 || len(data.FmtCode) > 32 || len(data.DtlCode) > 32 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//统计主表对应子表中有多少条数据
	count, err := querySubCount(data)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error()=="数据为空"{
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return
	}

	//数据不存在
	if count == 0 && err == nil {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return
	}

	//删除数据
	if err := subDelMethod(data); err != nil {
		fmt.Println(err.Error())
		if err.Error()=="数据不存在"{
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorDelErrId
		response.Msg = common.ErrorDelErrMsg
		return
	}

	//返回给前端删除剩余的条数
	type num struct {
		DelCount string `json:"delCount"`
	}

	//成功
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKDeleteMsg
	//如果删除成功当前这条数据则减去当前删除的数据—1
	response.Data = num{DelCount: fmt.Sprint(count - 1)}
}
