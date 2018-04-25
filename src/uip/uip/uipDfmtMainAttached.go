package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"errors"

	"uip/attached"
	"uip/common"
)

/*
	主表添加方法
*/
func mainAddMehod(data attached.UipDfmtMain) error {
	fmt.Println("***mainAddMehod***")
	//判断是否存在
	isExists, err := frame.DB.ExistsDo("exists", "uipDfmtMain"+data.GroupId+data.FmtCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	} else {
		if isExists[0] == "ok" && isExists[1] == "1" {
			return errors.New("数据已存在:" + err.Error())
		}
	}

	convertjsondata, _ := json.Marshal(data)
	fmt.Println("converdata:", string(convertjsondata))
	//向数据库插入数据
	if err := frame.DB.Create("uipDfmtMain", "uipDfmtMain"+data.GroupId+data.FmtCode, string(convertjsondata)); err != nil {
		return errors.New("添加错误:" + err.Error())
	}
	return nil
}

/*
	主表修改方法
*/
func mainUpdateMehod(data attached.UipDfmtMain) error {
	fmt.Println("***mainUpdateMehod***")
	//判断是否存在
	resultValue, err := frame.DB.RetriveOne("uipDfmtMain"+data.GroupId+data.FmtCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	}
	//存在状态
	if resultValue!=""&& err==nil {
		var mainFmt attached.UipDfmtMain
		if err!=json.Unmarshal([]byte(resultValue),&mainFmt){
			response.Code=common.ErrorJsonErrId
			response.Msg=common.ErrorJsonErrMsg
			return errors.New("json格式错误:"+err.Error())
		}
		//查询出来的时间在修改的时候写回去
		data.Ctime=mainFmt.Ctime

		convertjsondata, _ := json.Marshal(data)
		fmt.Println("converdata:", string(convertjsondata))

		//向数据库插入数据
		if err := frame.DB.Update("uipDfmtMain", "uipDfmtMain"+data.GroupId+data.FmtCode, string(convertjsondata)); err != nil {
			return errors.New("修改错误:" + err.Error())
		}
		return nil
	} else {
		//其他错误
		return errors.New("异常:" + err.Error())
	}
	return nil
}

/*
	主表删除方法
*/
func mainDelMethod(data attached.UipDfmtMain) error {
	fmt.Println("***mainUpdateMehod***")
	//判断是否存在
	isExists, err := frame.DB.ExistsDo("exists", "uipDfmtMain"+data.GroupId+data.FmtCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	}
	//存在状态
	if isExists[0] == "ok" && isExists[1] == "1" {
		convertjsondata, _ := json.Marshal(data)
		fmt.Println("converdata:", string(convertjsondata))
		//向数据库插入数据
		if err := frame.DB.Delete("uipDfmtMain", "uipDfmtMain"+data.GroupId+data.FmtCode); err != nil {
			return errors.New("删除错误:" + err.Error())
		}
		return nil
	} else {
		//不存在的状态
		return errors.New("数据不存在:" + err.Error())
	}
	return nil
}
