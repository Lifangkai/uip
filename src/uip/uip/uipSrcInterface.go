package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"uip/common"
	"uip/attached"
	"io/ioutil"
	"bytes"
)

type responseJson struct {
	 Code		string `json:"code"`
	 Msg		string `json:"msg"`
}
type responseJsonForUipSrcInterfacesGet struct {
	Code		string `json:"code"`
	Msg		string `json:"msg"`
	Data		[]attached.UipSrcInterface `json:"data"`
}
type Pse struct {
	Com string `json:"com"`
	Msg string `json:"msg"`
	Data []string `json:"data"`
}
type simpleFunc struct {
	FuncName	string `json:"funcName"`
	FuncCode	string `json:"funcCode"`
}
/*@Author:zengzeen
* @param:msrCountList struct
* @return:{code,msg,data}   json
 */
func UipSrcInterfaceHandlePostRequest( data attached.UipSrcInterface) {
	fmt.Println("this is UipSrcInterface post!")
	var err error
	var funcs []attached.UipInterFunc
	var  funcs2 []simpleFunc
	var insertSucc bool

	data.InteCode ,err = getSnsNew("UIPSRC1")
	data.Ctime,err = getSnsNew("dmm00001")
	emptyCheckResponse, isEmpty := UipSrcInterfaceParaFilter(data)
	if isEmpty == false {
		response.Msg = emptyCheckResponse.Msg
		response.Code = emptyCheckResponse.Code
		return
	}
	err = json.Unmarshal([]byte(data.FuncList), &funcs)
	if err != nil {
		response.Code = common.ErrorJsonFmtErrId
		response.Msg = common.ErrorJsonFmtErrMsg
		return
	}
	insertSucc = true
	for _,row :=range funcs {
		var a simpleFunc
		row.GroupId = data.GroupId
		row.Ctime ,err= getSnsNew("dmm00001")
		row.FuncCode ,err= getSnsNew("UIPFUNC1")
		if err != nil {
			insertSucc = false
			break
		}
		emptyCheckResponse, isEmpty := UipInterFuncParaFilter(row)
		if isEmpty == false {
			response.Code = emptyCheckResponse.Code
			response.Msg = emptyCheckResponse.Msg
			insertSucc = false
			break
		}
		isExists, errI := frame.DB.ExistsDo("exists", "uipInterFunc"+row.GroupId+row.FuncCode)
		if errI != nil {
			insertSucc = false
			break
		}
		if isExists[0] == "ok" && isExists[1] == "1" {
			insertSucc = false
			break
		}
		//向数据库插入新数据
		dataStr2, _ := json.Marshal(row)
		err =frame.DB.Create("uipInterFunc"+row.GroupId,"uipInterFunc"+row.GroupId+row.FuncCode, string(dataStr2))
		//截取func对象的funcCode和funcName字段，避免冗余
		a.FuncCode = row.FuncCode
		a.FuncName = row.FuncName
		funcs2 = append(funcs2,a)
		//异常处理
		if err != nil {
			insertSucc = false
			break
		}
	}
	if insertSucc == false{
		response.Code = common.ErrorInsertErrId
		response.Msg = common.ErrorInsertErrMsg
		return
	}

	//判断是否已经存在数据
	isExists, errI := frame.DB.ExistsDo("exists", "uipSrcInterface"+data.GroupId+data.InteCode)
	//异常处理
	if errI != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return
	}
	if isExists[0] == "ok" && isExists[1] == "1" {
		response.Code = common.ErrorDataExistsErrId
		response.Msg = common.ErrorDataExistsErrMsg
		return
	}

	//向数据库插入新数据
	funcStr, _ := json.Marshal(funcs2)
	data.FuncList = string(funcStr)
	fmt.Println("the funcList is:",data.FuncList)
	dataStr, _ := json.Marshal(data)
	err =frame.DB.Create("uipSrcInterface","uipSrcInterface"+data.GroupId+data.InteCode, string(dataStr))
	//异常处理
	if err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return
	}
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKInsertMsg
	response.Data = data
	return
}


/*
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipSrcInterfaceDelRequest( data attached.UipSrcInterface) {
	fmt.Println("this is UipSrcInterface Del method")
	var uip attached.UipSrcInterface
	var opens []attached.UipOpenInterface
	var errCode string
	var deleteSucc bool
	emptyCheckResponse, isEmpty := UipSrcInterfaceParaFilter(data)
	if isEmpty == false {
		response.Code = emptyCheckResponse.Code
		response.Msg = emptyCheckResponse.Msg
		return
	}
	OldData, err := frame.DB.RetriveOne("uipSrcInterface" + data.GroupId + data.InteCode)
	err = json.Unmarshal([]byte(OldData), &uip)
	if err != nil {
		response.Code = common.ErrorJsonFmtErrId
		response.Msg = common.ErrorJsonFmtErrMsg
		return
	}
	if len(uip.OpenList) == 0{
		response.Code = common.ErrorSrcInterfaceErrId
		response.Msg = common.ErrorSrcInterfaceMsg
		return
	}
	// 遍历uip_open_list表格
	err = json.Unmarshal([]byte(uip.OpenList), &opens)
	if err != nil {
		response.Code = common.ErrorJsonFmtErrId
		response.Msg = common.ErrorJsonFmtErrMsg
		return
	}
	deleteSucc = true
	for _,aOpen := range opens{
		err :=frame.DB.Delete("uipOpenInterface","uipSrcInterface" + aOpen.GroupId+data.InteCode)
		//异常处理
		if err != nil {
			errCode = aOpen.InteCode
			deleteSucc = false
			break
		} else {
			deleteSucc = true
			continue
		}
	}
	//返回给前端错误信息，和对应的编码
	if deleteSucc == false {
		response.Code = common.ErrorOpenInterDeleteErrId
		response.Msg = common.ErrorOpenInterDeleteMsg + " 错误的开放接口编码是:"+errCode
		return
	}

	if err == nil && OldData == "" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return
	} else if OldData == "" && err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return
	} else {
		 err :=frame.DB.Delete("uipSrcInterface","uipSrcInterface" + data.GroupId+data.InteCode)
		//异常处理
		if err != nil {
			response.Code = common.ErrorSystemErrId
			response.Msg = common.ErrorSystemErrMsg
			return
		} else {
			response.Code = common.ErrorOKId
			response.Msg = common.ErrorOKDeleteMsg
			return
		}
	}
}

/**
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipSrcInterfaceHandlePutRequest(data attached.UipSrcInterface) {
	fmt.Println("this is UipSrcInterface Put method")
	emptyCheckResponse, isEmpty := UipSrcInterfaceParaFilter(data)
	if isEmpty == false {
		response.Code = emptyCheckResponse.Code
		response.Msg = emptyCheckResponse.Msg
		return
	}
	//判断是否已经存在数据
	OldData, err := frame.DB.RetriveOne("uipSrcInterface" + data.GroupId + data.InteCode)
	if err == nil && OldData == "" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return
	}
	if OldData == "" && err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return
	} else {
		//向数据库插入新数据
		dataStr, _ := json.Marshal(data)
		err := frame.DB.Update("uipSrcInterface","uipSrcInterface"+data.GroupId+data.InteCode, string(dataStr))
		//异常处理
		if err != nil {
			response.Code = common.ErrorSystemErrId
			response.Msg = common.ErrorSystemErrMsg
			return
		}
		response.Code = common.ErrorOKId
		response.Msg = common.ErrorOKModifyMsg
		response.Data = data
		return
	}
}

/** GET方法 com=SN 根据groupId和sn查询数据
*	@Author zengzeen
*	@param  groupId 组织标识
*	@param  key 前来传过来用于查询的sn，即序列号
*   @return {code,msg,data} json
 */
func UipSrcInterfaceGetRequest( r *http.Request ) {
	fmt.Println("this is UipSrcInterface SNGet method")
	var groupId,inteCode string
	groupId = r.Form["groupId"][0]
	inteCode = r.Form["inteCode"][0]

	if len(groupId)==0 || len(inteCode)==0 {
		response.Code = common.ErrorDataNullId
		response.Msg = common.ErrorDataNullMsg
		return
	}else if len(groupId)!=8 {
		response.Code = common.ErrorGroupIdLenId
		response.Msg = common.ErrorGroupIdLenMsg
		return
	}else if len(inteCode)>32 {
		response.Code = common.ErrorInterLenId
		response.Msg = common.ErrorInterLenMsg
		return
	}

	//根据主键搜索
	result, err :=frame.DB.RetriveOne("uipSrcInterface"+groupId+inteCode)
	//fmt.Println("result.(type): ", reflect.TypeOf(result))
	if err == nil && result == "" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return
	} else if result == "" && err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return
	} else {
		response.Code = common.ErrorOKId
		response.Msg = common.ErrorOKSelectMsg
		var uip attached.UipSrcInterface

		err := json.Unmarshal([]byte(result), &uip)
		if err != nil {
			var response responseJson
			response.Code = common.ErrorJsonFmtErrId
			response.Msg = common.ErrorJsonFmtErrMsg
			return
		}
		response.Data = uip
		response.Code = common.ErrorOKId
		response.Msg = common.ErrorOKGetMsg
		return
	}
}

/**
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipSrcInterfacePseHandlePutRequest( r *http.Request) {
	fmt.Println("this is UipSrcInterface pse method")
	var keysGet Pse
	var resopnseData responseJsonForUipSrcInterfacesGet
	groupId := r.Form["groupId"][0]
	fettext := r.Form["fettext"][0]
	requestStr := "  {\"com\":\"search\",\"data\":{\"tableName\":\"uipSrcInterface"+groupId+"\",\"keyWords\":\""+fettext+"\"}}"
	var valstr=[]byte(requestStr)
	resp, err := http.Post("http://172.16.0.14:31007/pse",
		"application/json;charset=utf-8", bytes.NewBuffer(valstr))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	data2,err2:= ioutil.ReadAll(resp.Body)
	fmt.Println("the err when we read the respons's body is :",err2)
	err =json.Unmarshal([]byte(data2),&keysGet)
	fmt.Println("the keys is :",keysGet)
	resp.Body.Close()
	if err !=nil{
		fmt.Println("the json umarshal to keys[] is faild, the err is:",err)
		return
	}
	keys := keysGet.Data
	for _, keyword := range keys {
		var res attached.UipSrcInterface
		childRes := getlistOne(keyword)
		fmt.Println(childRes)
		if len(childRes.Data) != 0 {
			err := json.Unmarshal([]byte(childRes.Data), &res)
			//异常处理
			if err != nil || res.GroupId != groupId{
				fmt.Println("in pse method, some err occured:",err)
				continue
			}
			resopnseData.Data = append(resopnseData.Data, res)
		}
	}
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKGetMsg
	response.Data = resopnseData.Data
	return
}
