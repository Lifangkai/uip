package main

import (
	"fmt"
	"net/http"

	"uip/attached"
	"uip/common"
	"encoding/json"
)

/*
	添加数据方法
*/
func interfaceOpenInsertMethod(data attached.UipOpenInterface) {
	fmt.Println("***interfaceOpenInsertMethod***")

	//校验数据
	if ok := attached.OpenInterfaceFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//添加数据
	if err := openInterfaceAddMehod(data); err != nil {
		fmt.Println(err.Error())
		if err.Error() == "数据已存在" {
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
func interfaceOpenUpdateMethod(data attached.UipOpenInterface) {
	fmt.Println("***interfaceOpenUpdateMethod***")

	//校验数据
	if ok := attached.OpenInterfaceFilter(data); !ok {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//修改
	if err := openInterfaceUpdateMehod(data); err != nil {
		fmt.Println(err.Error())
		if err.Error() == "数据不存在" {
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
	删除数据方法，同时删除源接口中的对应本条的openlist字段
*/
func interfaceOpenDeleteMethod(data attached.UipOpenInterface) {
	fmt.Println("***interfaceOpenDeleteMethod***")

	//数据过滤
	if data.GroupId == "" || data.InteCode == "" || len(data.GroupId) != 8 || len(data.InteCode) > 32 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//先查询当前openInterface表数据是否存在
	resultInterface, err := openInterfaceQuery(data.GroupId + data.InteCode)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "数据不存在" {
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorQueryErrId
		response.Msg = common.ErrorQueryErrMsg
		return
	}

	//查询openInterface对应的srcInterface是否存在(开放接口sinteCode对应源接口中的inteCode)
	resultSrcInterface, err := querySrcInterfaceMethod(resultInterface.GroupId + resultInterface.SinteCode)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "数据不存在" {
			response.Code = common.ErrorSrcInterfaceDataIsNilId
			response.Msg = common.ErrorSrcInterfaceDataIsNilMsg
			return
		}
		response.Code = common.ErrorQueryErrId
		response.Msg = common.ErrorQueryErrMsg
		return
	}

	//序列化SrcInterFace表中OpenList字段成员所用结构体
	type openlist struct {
		InteCode string `json:"inteCode"`
		InteName string `json:"inteName"`
		InteType string `json:"inteType"`
	}

	//序列化所用对象
	var ol []openlist

	//删除srcInterface表中成员字段openList中对应OpenInterface表的信息(本质更新)
	if err := json.Unmarshal([]byte(resultSrcInterface.OpenList), &ol); err != nil {
		response.Code = common.ErrorJsonErrId
		response.Msg = common.ErrorJsonErrMsg
		return
	}

	boo := false
	for i, v := range ol {
		//判断当前openInterface的信息是否在这条OpenList当中
		if v.InteCode == resultInterface.InteCode {
			boo = true

			/*
				example:delete '3'
				array:=[]int{1,2,3,4,5}
				i==2
				array[i:]==[3,4,5]
				array[i+1:]==[4,5]
				copy()==array==[1,2,4,5,5]
				finalValue:[:len(array)-1]==[1,2,4,5]
			*/

			//删除当前存在的数据
			copy(ol[i:], ol[i+1:])
			break
		}
	}

	//不存在状态
	if boo == false {
		response.Code = common.ErrorInterfaceNotExsitId
		response.Msg = common.ErrorInterfaceNotExsitMsg
		return
	}

	//把数据写回去(修改)
	updateValue := ol[:len(ol)-1] //取出最后切片中最后一个元素

	//中间变量
	var src attached.UipSrcInterface
	src = resultSrcInterface

	//转成json
	convertData, _ := json.Marshal(updateValue)

	//写回
	src.OpenList = string(convertData)

	//修改
	if err := updateSrcInterfaceMethod(resultInterface.GroupId+resultInterface.SinteCode, src); err != nil {
		response.Code = common.ErrorDelSrcInterfaceErrId
		response.Msg = common.ErrorDelSrcInterfaceErrMsg + data.InteCode //提示错误并带上开放接口的编码
		return
	}

	//删除openInterface
	if err := openInterfaceDelMethod(data); err != nil {
		fmt.Println(err)
		if err.Error() == "数据不存在" {
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorOpenInterfaceDataIsNilId
		response.Msg = common.ErrorOpenInterfaceDataIsNilMsg
		return
	}

	//成功
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKDeleteMsg
}

/*
	开放接口查询单条
*/
func interfaceOpenQueryOneMethod(r *http.Request) {
	fmt.Println("***interfaceOpenQueryOneMethod***")

	//接受数据并校验
	if r.Form["groupId"][0] == "" || len(r.Form["groupId"]) != 1 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	if r.Form["inteCode"][0] == "" || len(r.Form["inteCode"]) != 1 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//查询的条件
	key := r.Form["groupId"][0] + r.Form["inteCode"][0]

	//查询
	result, err := openInterfaceQuery(key)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "数据不存在" {
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorQueryErrId
		response.Msg = common.ErrorQueryErrMsg
		return
	}

	//返回
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKSelectMsg
	response.Data = result
}

/*
	查询主表与子表的数据方法
*/
func interfaceOpenQueryManyMethod(r *http.Request) {
	fmt.Println("***interfaceOpenQueryManyMethod***")

	//接受数据并校验
	if r.Form["groupId"][0] == "" || len(r.Form["groupId"]) != 1 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	if r.Form["ftText"][0] == "" || len(r.Form["ftText"]) != 1 {
		response.Code = common.CheckErrorId
		response.Msg = common.CheckErrorMsg
		return
	}

	//查询的条件
	//key := r.Form["groupId"][0] + r.Form["ftText"][0]

	queryData, err := queryOpenInterfaceAll(r.Form["groupId"][0], r.Form["ftText"][0])
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "数据不存在" {
			response.Code = common.ErrorDataNotExistsErrId
			response.Msg = common.ErrorDataNotExistsMsg
			return
		}
		response.Code = common.ErrorQueryErrId
		response.Msg = common.ErrorQueryErrMsg
		return
	}

	//返回
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKSelectMsg
	response.Data = queryData
}
