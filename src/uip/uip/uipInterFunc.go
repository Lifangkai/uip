package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"uip/common"
	"io/ioutil"
	"bytes"
	"uip/attached"
)

/**
* @param:uipInterFunc struct
* @return:{code,msg,data}   json
 */
func UipInterFuncHandlePostRequest( data attached.UipInterFunc) {
	common.Frame.Log.Write("in UipInterFuncHandlePostRequest（）")

	var err error
	data.FuncCode , err = getSnsNew(data.GroupId ,"UIPFUNC1")
	data.Ctime, err = getSnsNew(data.GroupId ,"dmm00001")
	//data.ReqFmtCode,err = getSnsNew("reqFmtC")  //前端传入
	//data.ResFmtCode,err = getSnsNew("resFmtC")  //前端传入
	
	emptyCheckResponse, isEmpty := UipInterFuncParaFilter(data)
	if isEmpty == false {
		response.Code = emptyCheckResponse.Code
		response.Msg = emptyCheckResponse.Msg
		return
	}
	
	//判断是否已经存在数据
	isExists, errI := common.Frame.DB.ExistsDo("exists", "uipInterFunc"+data.GroupId+data.InteCode+data.FuncCode)
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
	dataStr, _ := json.Marshal(data)
	err =common.Frame.DB.Create("uipInterFunc"+data.GroupId,"uipInterFunc"+data.GroupId+data.InteCode+data.FuncCode, string(dataStr))
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

/**
*/
func UipInterFuncsHandlePostRequest(datas []attached.UipInterFunc) {
	common.Frame.Log.Write("this is UipInterFunc post!")
	var insertSucc  bool
	var err error
	insertSucc = true
	
	for i,data:= range datas{
		data.FuncCode ,err = getSnsNew(data.GroupId ,"UIPFUNC1")
		data.Ctime,err = getSnsNew(data.GroupId ,"dmm00001")
		data.ReqFmtCode,err = getSnsNew(data.GroupId ,"reqFmtC")
		data.ResFmtCode,err = getSnsNew(data.GroupId ,"resFmtC")
		datas[i].FuncCode = data.FuncCode
		datas[i].Ctime = data.Ctime
		emptyCheckResponse,isEmpty := UipInterFuncParaFilter(data)
		
		if isEmpty == false {
			response.Code = emptyCheckResponse.Code
			response.Msg = emptyCheckResponse.Msg
			return

		}
		
		//判断是否已经存在数据
		isExists, errI := common.Frame.DB.ExistsDo("exists", "uipInterFunc"+data.GroupId+data.InteCode+data.FuncCode)
		//异常处理
		if errI != nil {
			insertSucc = false
			break
		}
		
		if isExists[0] == "ok" && isExists[1] == "1" {
			insertSucc = false
			break
		}

		//向数据库插入新数据
		dataStr, _ := json.Marshal(data)
		err =common.Frame.DB.Create("uipSrcInterface"+data.GroupId,"uipInterFunc"+data.GroupId+data.InteCode+data.FuncCode, string(dataStr))
		//异常处理
		if err != nil {
			insertSucc = false
			break

		}
	}
	
	if insertSucc{
		response.Code = common.ErrorOKId
		response.Msg = common.ErrorOKInsertMsg
		return
	}else {
		response.Code = common.InsertErrorId
		response.Msg = common.InsertErrorMsg
		response.Data = datas
		return
	}

}
/**
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */
func UipInterFuncvDelRequest( data attached.UipInterFunc) {
	common.Frame.Log.Write("this is UipInterFunc Del method")
	dataStr, _ := json.Marshal(data)
	var uip attached.UipInterFunc
	
	err := json.Unmarshal([]byte(dataStr), &uip)
	if err != nil {
		response.Code = common.ErrorJsonFmtErrId
		response.Msg = common.ErrorJsonFmtErrMsg
		return
	}

	//emptyCheckResponse,isEmpty := UipInterFuncParaFilter(data)
	//if isEmpty == false {
	//	response.Code = emptyCheckResponse.Code
	//	response.Msg = emptyCheckResponse.Msg
	//	return
	//
	//}
	
	if len(data.GroupId)==0 || len(data.InteCode)== 0 || len(data.FuncCode) == 0 {
		response.Code = common.ErrorDataNullId
		response.Msg = common.ErrorDataNullMsg
		return
	}
	
	OldData, err := common.Frame.DB.RetriveOne("uipInterFunc" + data.GroupId +data.InteCode+ data.FuncCode)
	if err == nil && OldData == "" {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return
	} else if OldData == "" && err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
		return

	} else {
		err :=common.Frame.DB.Delete("uipInterFunc"+data.GroupId,"uipInterFunc" + data.GroupId+data.InteCode+data.FuncCode)
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
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */
func UipInterFuncHandlePutRequest( data attached.UipInterFunc) {
	common.Frame.Log.Write("this is uipInterFunc Put method")
	
	emptyCheckResponse,isEmpty := UipInterFuncParaFilter(data)
	if isEmpty == false {
		response.Code = emptyCheckResponse.Code
		response.Msg = emptyCheckResponse.Msg
		return
	}
	
	//判断是否已经存在数据
	OldData, err := common.Frame.DB.RetriveOne("uipInterFunc" + data.GroupId +data.InteCode+ data.FuncCode)
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
		err := common.Frame.DB.Update("uipInterFunc"+data.GroupId,"uipInterFunc"+data.GroupId+data.InteCode+data.FuncCode, string(dataStr))
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
*	@param  groupId 组织标识
*	@param  funcCode 前来传过来用于查询的sn，即序列号
*   @return {code,msg,data} json
 */
func UipInterFuncGetRequest(r *http.Request ) {
	common.Frame.Log.Write("this is uipInterFunc SNGet method")
	var groupId,funcCode,inteCode string

	groupId = r.Form["groupId"][0]
	inteCode = r.Form["inteCode"][0]
	funcCode = r.Form["funcCode"][0]

	if len(groupId)==0 || len(funcCode)==0 {
		response.Code = common.ErrorDataNullId
		response.Msg = common.ErrorDataNullMsg
		return
	}else if len(groupId)!=8 {
		response.Code = common.ErrorGroupIdLenId
		response.Msg = common.ErrorGroupIdLenMsg
		return
	}else if len(funcCode)>32 {
		response.Code = common.ErrorFuncLenId
		response.Msg = common.ErrorFuncLenMsg
		return
	}

	//根据主键搜索
	result, err :=common.Frame.DB.RetriveOne("uipInterFunc"+groupId+inteCode+funcCode)
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
		var uip attached.UipInterFunc

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
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
*/
func UipInterFuncPseHandlePutRequest( r *http.Request) {
	common.Frame.Log.Write("this is uipInterFunc pse method")
	var datas []attached.UipInterFunc
	var keysGet Pse
	groupId := r.Form["groupId"][0]
	settext := r.Form["settext"][0]
	requestStr := " {\"com\":\"search\",\"data\":{\"tableName\":\"uipInterFunc"+groupId+"\",\"keyWords\":\""+settext+"\"}}"
	var valstr=[]byte(requestStr)
	
	resp, err := http.Post("http://"+ common.Frame.PseIp + ":"+ common.Frame.PsePort +"/pse",
		"application/json;charset=utf-8", bytes.NewBuffer(valstr))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	
	data2,err2:= ioutil.ReadAll(resp.Body)
	fmt.Println("the err when we read the respons's body is :",err2)
	err =json.Unmarshal([]byte(data2),&keysGet)
	resp.Body.Close()
	if err !=nil{
		fmt.Println("the json umarshal to keys[] is faild, the err is:",err)
		return
	}
	
	keys := keysGet.Data
	for _, keyword := range keys {
		var res attached.UipInterFunc
		childRes := getlistOne(keyword)
		fmt.Println(childRes)
		if len(childRes.Data) != 0 {
			err := json.Unmarshal([]byte(childRes.Data), &res)
			//异常处理
			if err != nil || res.GroupId != groupId{
				fmt.Println("in pse method, some err occured:",err)
				continue
			}
			datas = append(datas, res)
		}
	}
	
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKGetMsg
	response.Data = datas
	
	return
}

/**
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
*/
func UipInterFuncvBatchQueryRequest( datas attached.BatchResQuest) {
	common.Frame.Log.Write("this is uipInterFunc batchQuery method")
	var resDatas []attached.UipInterFunc
	var querySucc bool
	groupId := datas.GroupId
	querySucc = true
	
	for _,data := range  datas.Data{
		result, err :=common.Frame.DB.RetriveOne("uipInterFunc"+groupId+data)
		if err == nil && result == "" {
			querySucc = false
			break
		} else if result == "" && err != nil {
			querySucc = false
			break
		} else {
			var uip attached.UipInterFunc
			err := json.Unmarshal([]byte(result), &uip)
			if err != nil {
				querySucc = false
				break
			}
			resDatas = append(resDatas,uip)
		}
	}
	
	if len(resDatas) == 0 {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
		return
	}
	
	if querySucc {
		response.Code = common.ErrorOKId
		response.Msg = common.ErrorOKGetMsg
		response.Data = resDatas
		return
	}else {
		response.Code = common.ErrorQueryId
		response.Msg = common.ErrorQueryMsg
		return
	}
}


/** 
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
*/
func UipInterFuncQuryAlltRequest( r *http.Request) {
	common.Frame.Log.Write("this is uipInterFunc queryall  method")
	var datas []attached.UipInterFunc
	var uip attached.UipInterFunc
	groupId := r.Form["groupId"][0]
	inteCode := r.Form["inteCode"][0]
	
	results, err := common.Frame.DB.Retrive("scan", "uipInterFunc"+groupId+inteCode+"0", "uipInterFunc"+groupId+inteCode+"~", 999999)
	if err != nil {
		response.Code = common.ErrorSystemErrId
		response.Msg = common.ErrorSystemErrMsg
	}
	if len(results) < 0 {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
	}

	for i := 2; i < len(results); i += 2 {
		if err:=json.Unmarshal([]byte(results[i]), &uip); err!= nil{
			continue
		}
		datas= append(datas,uip)
	}
	
	if len(datas) < 0 {
		response.Code = common.ErrorDataNotExistsErrId
		response.Msg = common.ErrorDataNotExistsMsg
	}
	
	response.Code = common.ErrorOKId
	response.Msg = common.ErrorOKGetMsg
	response.Data = datas
	
	return
}
