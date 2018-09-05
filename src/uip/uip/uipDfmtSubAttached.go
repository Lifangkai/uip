package main

import (
	"fmt"
	"time"
	"errors"
	"encoding/json"

	"uip/attached"
	"uip/common"
)

/**
	子表添加方法
*/
func subAddMehod(data attached.UipDfmtSub) error {
	frame.Log.Write("***subAddMehod***")
	//判断是否存在
	isExists, err := frame.DB.ExistsDo("exists", "uipDfmtSub"+data.GroupId+data.FmtCode+data.DtlCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	} else {
		if isExists[0] == "ok" && isExists[1] == "1" {
			return errors.New("数据已存在")
		}
	}

	//获取当前时间
	t := time.Now().Format("20060102150405")
	data.Ctime = t

	convertJsonData, _ := json.Marshal(data)
	fmt.Println("converdata:", string(convertJsonData))
	//向数据库插入数据
	if err := frame.DB.Create("uipDfmtSub", "uipDfmtSub"+data.GroupId+data.FmtCode+data.DtlCode, string(convertJsonData)); err != nil {
		return errors.New("添加错误:" + err.Error())
	}
	return nil
}

/**
	子表修改方法
*/
func subUpdateMehod(data attached.UipDfmtSub) error {
	frame.Log.Write("***subUpdateMehod***")
	//判断是否存在
	resultValue, err := frame.DB.RetriveOne("uipDfmtSub" + data.GroupId + data.FmtCode + data.DtlCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	}
	//存在状态
	if resultValue != "" && err == nil {
		var subFmt attached.UipDfmtSub
		if err != json.Unmarshal([]byte(resultValue), &subFmt) {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return errors.New("json格式错误:" + err.Error())
		}
		//查询出来的时间在修改的时候写回去
		data.Ctime = subFmt.Ctime

		convertJsonData, _ := json.Marshal(data)
		fmt.Println("converdata:", string(convertJsonData))

		//向数据库插入数据
		if err := frame.DB.Update("uipDfmtSub", "uipDfmtSub"+data.GroupId+data.FmtCode+data.DtlCode, string(convertJsonData)); err != nil {
			return errors.New("修改错误:" + err.Error())
		}
		return nil
	} else {
		//其他错误
		return errors.New("数据不存在")
	}
	return nil
}

/**
	子表删除方法
*/
func subDelMethod(data attached.UipDfmtSub) error {
	frame.Log.Write("***subDelMethod***")
	//判断是否存在
	isExists, err := frame.DB.ExistsDo("exists", "uipDfmtSub"+data.GroupId+data.FmtCode+data.DtlCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	}
	//存在状态
	if isExists[0] == "ok" && isExists[1] == "1" {
		convertJsonData, _ := json.Marshal(data)
		fmt.Println("converdata:", string(convertJsonData))
		//向数据库插入数据
		if err := frame.DB.Delete("uipDfmtSub", "uipDfmtSub"+data.GroupId+data.FmtCode+data.DtlCode); err != nil {
			return errors.New("删除错误")
		}
		return nil
	} else {
		//不存在的状态
		return errors.New("数据不存在")
	}
	return nil
}

//子表查询所有用于统计条数
func querySubCount(data attached.UipDfmtSub) (count int, err error) {
	fmt.Println("***querySubCount***")

	key := "uipDfmtSub" + data.GroupId + data.FmtCode

	//查询所有
	results, err := frame.DB.Retrive("scan", key+"0", key+"~", 999999999)
	//异常处理
	if err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return 0, errors.New("查询所有错误:" + err.Error())
	}

	//成功但是并没有数据
	if len(results) == 1 && results[0] == "ok" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return 0, errors.New("数据为空")
	} else {
		long := len(results)
		//计次
		var c int
		for i := 2; i < long; i += 2 {
			c++
		}
		return c, nil
	}
}

//子表查询所有
func querySubAllCondtion(key string) ([]attached.UipDfmtSub,error) {
	frame.Log.Write("***querySubAllCondtion***")

	//返回的实体
	var subArray []attached.UipDfmtSub

	//key条件
	Key := "uipDfmtSub" + key

	//查询所有
	results, err := frame.DB.Retrive("scan", Key+"0", Key+"~", 999999999)
	//异常处理
	if err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return nil, errors.New("查询所有错误:" + err.Error())
	}

	//成功但是并没有数据
	if len(results) == 1 && results[0] == "ok" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return nil,errors.New("数据为空" )
	} else {
		//反序列化的实体
		var subBean attached.UipDfmtSub

		long := len(results)
		for i := 2; i < long; i += 2 {
			if err := json.Unmarshal([]byte(results[i]), &subBean); err != nil {
				response.Code = common.ErrorJsonErrId
				response.Msg = common.ErrorJsonErrMsg
				return nil, errors.New("json反序列化错误:" + err.Error())
			}
			//当前序列化成功的对象追加
			subArray = append(subArray, subBean)
		}

		if subArray == nil {
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return nil, errors.New("数据为空")
		}
		return subArray, nil
	}
}
