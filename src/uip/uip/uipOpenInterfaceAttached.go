package main

import (
	"encoding/json"
	"errors"
	"uip/attached"
	"fmt"
	"time"
	"uip/common"
)

/*
	开放接口单条添加方法
*/
func openInterfaceAddMehod(data attached.UipOpenInterface) error {
	fmt.Println("***openInterfaceAddMehod***")

	//判断是否存在
	isExists, err := frame.DB.ExistsDo("exists", "uipOpenInterface"+data.GroupId+data.InteCode)
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
	if err := frame.DB.Create("uipOpenInterface", "uipOpenInterface"+data.GroupId+data.InteCode, string(convertJsonData)); err != nil {
		return errors.New("添加错误:" + err.Error())
	}
	return nil
}

/*
	开放接口单条修改方法
*/
func openInterfaceUpdateMehod(data attached.UipOpenInterface) error {
	fmt.Println("***openInterfaceUpdateMehod***")

	//判断是否存在
	resultValue, err := frame.DB.RetriveOne("uipOpenInterface" + data.GroupId + data.InteCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	}

	//存在状态
	if resultValue != "" && err == nil {
		var openInter attached.UipOpenInterface
		if err != json.Unmarshal([]byte(resultValue), &openInter) {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return errors.New("json格式错误:" + err.Error())
		}
		//查询出来的时间在修改的时候写回去
		data.Ctime = openInter.Ctime

		convertJsonData, _ := json.Marshal(data)
		fmt.Println("converdata:", string(convertJsonData))

		//向数据库插入数据
		if err := frame.DB.Update("uipOpenInterface", "uipOpenInterface"+data.GroupId+data.InteCode, string(convertJsonData)); err != nil {
			return errors.New("修改错误:" + err.Error())
		}
		return nil
	} else {
		//其他错误
		return errors.New("数据不存在")
	}
	return nil
}

/*
	开放接口单条删除方法
*/
func openInterfaceDelMethod(data attached.UipOpenInterface) error {
	fmt.Println("***openInterfaceDelMethod***")

	//判断是否存在
	isExists, err := frame.DB.ExistsDo("exists", "uipOpenInterface"+data.GroupId+data.InteCode)
	//异常处理
	if err != nil {
		return errors.New("判断是否存在出错:" + err.Error())
	}

	//存在状态
	if isExists[0] == "ok" && isExists[1] == "1" {
		convertJsonData, _ := json.Marshal(data)
		fmt.Println("converdata:", string(convertJsonData))
		//向数据库插入数据
		if err := frame.DB.Delete("uipOpenInterface", "uipOpenInterface"+data.GroupId+data.InteCode); err != nil {
			return errors.New("删除错误")
		}
		return nil
	} else {
		//不存在的状态
		return errors.New("数据不存在")
	}
	return nil
}

/*
	开放接口单条查询方法
*/
func openInterfaceQuery(key string) (attached.UipOpenInterface, error) {
	fmt.Println("***openInterfaceQuery***")

	//用于返回的值
	var openInterFace attached.UipOpenInterface

	//条件
	Key := "uipOpenInterface" + key

	//查询
	result, err := frame.DB.RetriveOne(Key)

	//异常处理
	if err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return openInterFace, errors.New("查询错误:" + err.Error())
	}

	//成功但是并没有数据
	if result == "" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return openInterFace, errors.New("数据不存在")
	}

	//序列化成结构体
	if err := json.Unmarshal([]byte(result), &openInterFace); err != nil {
		response.Code = common.ErrorJsonErrId
		response.Msg = common.ErrorJsonErrMsg
		return openInterFace, errors.New("json序列化错误:" + err.Error())
	}
	return openInterFace, nil
}

/*
	开放接口通过搜索引擎(pse)查询前20条方法
*/
func queryOpenInterfaceAll(groupId, condition string) ([]attached.UipOpenInterface, error) {
	fmt.Println("***queryOpenInterfaceAll***")

	//返回的实体
	var openInterfaceArray []attached.UipOpenInterface

	//key条件
	//Key := "uipOpenInterface" + key

	//调用搜索引擎
	getKeysByPse, err := attached.GetPse("uipOpenInterface", groupId, condition)
	if err != nil {
		fmt.Println("pse错误:" + err.Error())
		return nil, errors.New(err.Error())
	}

	//用于取前20个key
	var newKey []string
	newKey = getKeysByPse

	//只查询前20所以进行判断,如果大于20则取前二十,反之
	if len(getKeysByPse) > 20 {
		//覆盖
		newKey = getKeysByPse[:20]
	}

	//根据pse查询出来的key数组进行查询
	for _, v := range newKey {
		result, err := openInterfaceQuery(v)
		if err != nil {
			fmt.Println("查询当前失败:" + err.Error())
			continue
		}

		//把当前查询出来的值进行追加
		openInterfaceArray = append(openInterfaceArray, result)
	}

	//为空判断
	if openInterfaceArray == nil {
		return nil, errors.New("数据不存在")
	}

	//查询出来的值进行返回
	return openInterfaceArray, nil
}

/*
	查询源接口单条数据方法
*/

func querySrcInterfaceMethod(key string) (attached.UipSrcInterface, error) {
	fmt.Println("***querySrcInterfaceMethod***")

	//用于返回的值
	var srcInterFace attached.UipSrcInterface

	//条件
	Key := "uipSrcInterface" + key

	//查询
	result, err := frame.DB.RetriveOne(Key)

	//异常处理
	if err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return srcInterFace, errors.New("查询错误:" + err.Error())
	}

	//成功但是并没有数据
	if result == "" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return srcInterFace, errors.New("数据不存在")
	}

	//序列化成结构体
	if err := json.Unmarshal([]byte(result), &srcInterFace); err != nil {
		response.Code = common.ErrorJsonErrId
		response.Msg = common.ErrorJsonErrMsg
		return srcInterFace, errors.New("json序列化错误:" + err.Error())
	}
	return srcInterFace, nil
}

/*
	接口单条数据修改方法
*/
func updateSrcInterfaceMethod(key string, srcInterface attached.UipSrcInterface) error {
	fmt.Println("***updateSrcInterfaceMethod***")

	//条件
	Key := "uipSrcInterface" + key

	//转成json
	convertData, _ := json.Marshal(srcInterface)

	//修改
	if err := frame.DB.Update("uipSrcInterface", Key, string(convertData)); err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return errors.New("修改错误:" + err.Error())
	}

	//成功返回
	return nil
}
