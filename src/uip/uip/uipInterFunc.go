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


type responseJsonForUipInterFuncGet struct {
	Code		string `json:"code"`
	Msg			string `json:"msg"`
	Data		attached.UipInterFunc `json:"data"`
}
type responseJsonForUipInterFuncsGet struct {
	Code		string `json:"code"`
	Msg			string `json:"msg"`
	Data		[]attached.UipInterFunc `json:"data"`
}


/*@Author:zengzeen
* @param:uipInterFunc struct
* @return:{code,msg,data}   json
 */
func UipInterFuncHandlePostRequest(w http.ResponseWriter, r *http.Request, data attached.UipInterFunc) {
	fmt.Println("this is UipInterFunc post!")
	var responseInfo responseJson
	var responseStr []byte

	//获取输入数据
	dataStr, _ := json.Marshal(data)
	fmt.Println("dataStr=[", dataStr, "]")
	var uip attached.UipInterFunc
	err := json.Unmarshal([]byte(dataStr), &uip)

	if err != nil {
		responseInfo.Code = common.ErrorJsonFmtErrId
		responseInfo.Msg = common.ErrorJsonFmtErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	data.FuncCode ,err = getSnsNew("UIPFUNC1")
	data.Ctime,err = getSnsNew("dmm00001")
	emptyCheckResponse, isEmpty := UipInterFuncParaFilter(uip)
	if isEmpty == false {
		responseStr, _ = json.Marshal(emptyCheckResponse)
		w.Write([]byte(responseStr))
		return
	}
	//判断是否已经存在数据
	isExists, errI := frame.DB.ExistsDo("exists", "uipInterFunc"+data.GroupId+data.FuncCode)
	//异常处理
	if errI != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	if isExists[0] == "ok" && isExists[1] == "1" {
		responseInfo.Code = common.ErrorDataExistsErrId
		responseInfo.Msg = common.ErrorDataExistsErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}

	//向数据库插入新数据
	err =frame.DB.Create("uipSrcInterface","uipSrcInterface"+data.GroupId+data.FuncCode, string(dataStr))
	//异常处理
	if err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return

	}
	responseInfo.Code = common.ErrorOKId
	responseInfo.Msg = common.ErrorOKInsertMsg
	responseStr, _ = json.Marshal(responseInfo)
	w.Write([]byte(responseStr))
	return
}

func UipInterFuncsHandlePostRequest(w http.ResponseWriter, r *http.Request, datas []attached.UipInterFunc) {
	fmt.Println("this is UipInterFunc post!")
	var responseInfo responseJson
	var responseStr []byte
	var insertSucc  bool
	for _,data:= range datas{
		//获取输入数据
		dataStr, _ := json.Marshal(data)
		fmt.Println("dataStr=[", dataStr, "]")
		var uip attached.UipInterFunc
		err := json.Unmarshal([]byte(dataStr), &uip)

		if err != nil {
			insertSucc = false
			break
		}
		data.FuncCode ,err = getSnsNew("UIPFUNC1")
		data.Ctime,err = getSnsNew("dmm00001")
		_, isEmpty := UipInterFuncParaFilter(uip)
		if isEmpty == false {
			insertSucc = false
			break
		}
		//判断是否已经存在数据
		isExists, errI := frame.DB.ExistsDo("exists", "uipInterFunc"+data.GroupId+data.FuncCode)
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
		err =frame.DB.Create("uipSrcInterface","uipInterFunc"+data.GroupId+data.FuncCode, string(dataStr))
		//异常处理
		if err != nil {
			insertSucc = false
			break

		}
	}
	if insertSucc{
		responseInfo.Code = common.ErrorOKId
		responseInfo.Msg = common.ErrorOKInsertMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}else {
		responseInfo.Code = common.InsertErrorId
		responseInfo.Msg = common.InsertErrorMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}

}
/*
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipInterFuncvDelRequest(w http.ResponseWriter, r *http.Request, data attached.UipInterFunc) {
	fmt.Println("this is UipInterFunc Del method")
	var responseInfo responseJson
	var responseStr []byte
	dataStr, _ := json.Marshal(data)
	var uip attached.UipInterFunc
	err := json.Unmarshal([]byte(dataStr), &uip)
	if err != nil {
		responseInfo.Code = common.ErrorJsonFmtErrId
		responseInfo.Msg = common.ErrorJsonFmtErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}

	emptyCheckResponse,isEmpty := UipInterFuncParaFilter(data)
	if isEmpty == false {
		responseStr, _ = json.Marshal(emptyCheckResponse)
		w.Write([]byte(responseStr))
		return

	}
	OldData, err := frame.DB.RetriveOne("uipInterFunc" + data.GroupId + data.FuncCode)
	if err == nil && OldData == "" {
		responseInfo.Code = common.ErrorDataNotExistsErrId
		responseInfo.Msg = common.ErrorDataNotExistsMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
	} else if OldData == "" && err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))

	} else {
		err :=frame.DB.Delete("uipInterFunc","uipInterFunc" + data.GroupId+data.FuncCode)
		//异常处理
		if err != nil {
			responseInfo.Code = common.ErrorSystemErrId
			responseInfo.Msg = common.ErrorSystemErrMsg
			responseStr, _ = json.Marshal(responseInfo)
			w.Write([]byte(responseStr))
		} else {
			responseInfo.Code = common.ErrorOKId
			responseInfo.Msg = common.ErrorOKDeleteMsg
			responseStr, _ = json.Marshal(responseInfo)
			w.Write([]byte(responseStr))
		}
	}
}

/**
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipInterFuncHandlePutRequest(w http.ResponseWriter, r *http.Request, data attached.UipInterFunc) {
	fmt.Println("this is uipInterFunc Put method")
	var responseInfo responseJson
	var responseStr []byte

	dataStr, _ := json.Marshal(data)
	var uip attached.UipInterFunc
	err := json.Unmarshal([]byte(dataStr), &uip)
	if err != nil {
		responseInfo.Code = common.ErrorJsonFmtErrId
		responseInfo.Msg = common.ErrorJsonFmtErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	emptyCheckResponse, isEmpty := UipInterFuncParaFilter(uip)
	if isEmpty == false {
		responseStr, _ = json.Marshal(emptyCheckResponse)
		w.Write([]byte(responseStr))
		return
	}
	//判断是否已经存在数据
	OldData, err := frame.DB.RetriveOne("uipInterFunc" + data.GroupId + data.FuncCode)
	if err == nil && OldData == "" {
		responseInfo.Code = common.ErrorDataNotExistsErrId
		responseInfo.Msg = common.ErrorDataNotExistsMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	if OldData == "" && err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	} else {
		fmt.Println("UipSrcInterface = [", uip, "]")
		//向数据库插入新数据
		err := frame.DB.Update("uipInterFunc","uipInterFunc"+data.GroupId+data.FuncCode, string(dataStr))
		//异常处理
		if err != nil {
			responseInfo.Code = common.ErrorSystemErrId
			responseInfo.Msg = common.ErrorSystemErrMsg
			responseStr, _ = json.Marshal(responseInfo)
			w.Write([]byte(responseStr))
		}
		responseInfo.Code = common.ErrorOKId
		responseInfo.Msg = common.ErrorOKModifyMsg
		responseStr, _ := json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
}

/** GET方法 com=SN 根据groupId和sn查询数据
*	@Author zengzeen
*	@param  groupId 组织标识
*	@param  funcCode 前来传过来用于查询的sn，即序列号
*   @return {code,msg,data} json
 */
func UipInterFuncGetRequest(w http.ResponseWriter, r *http.Request ) {
	fmt.Println("this is uipInterFunc SNGet method")
	var responseInfo responseJson
	var responseStr []byte
	var interFuncResponse responseJsonForUipInterFuncGet
	var groupId,funcCode string

	groupId = r.Form["groupId"][0]
	funcCode = r.Form["funcCode"][0]

	if len(groupId)==0 || len(funcCode)==0 {
		responseInfo.Code = common.ErrorDataNullId
		responseInfo.Msg = common.ErrorDataNullMsg
		return
	}else if len(groupId)!=8 {
		responseInfo.Code = common.ErrorGroupIdLenId
		responseInfo.Msg = common.ErrorGroupIdLenMsg
		return
	}else if len(funcCode)>32 {
		responseInfo.Code = common.ErrorFuncLenId
		responseInfo.Msg = common.ErrorFuncLenMsg
		return
	}

	//根据主键搜索
	result, err :=frame.DB.RetriveOne("uipInterFunc"+groupId+funcCode)
	if err == nil && result == "" {
		responseInfo.Code = common.ErrorDataNotExistsErrId
		responseInfo.Msg = common.ErrorDataNotExistsMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	} else if result == "" && err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	} else {
		responseInfo.Code = common.ErrorOKId
		responseInfo.Msg = common.ErrorOKSelectMsg
		var uip attached.UipInterFunc

		err := json.Unmarshal([]byte(result), &uip)
		if err != nil {
			var response responseJson
			response.Code = common.ErrorJsonFmtErrId
			response.Msg = common.ErrorJsonFmtErrMsg
			responseStr, _ = json.Marshal(response)
			w.Write([]byte(responseStr))
			return
		}
		interFuncResponse.Data = uip
		interFuncResponse.Code = common.ErrorOKId
		interFuncResponse.Msg = common.ErrorOKGetMsg
		responseStr, _ = json.Marshal(interFuncResponse)
		w.Write([]byte(responseStr))
	}
}

/**
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipInterFuncPseHandlePutRequest(w http.ResponseWriter, r *http.Request, data attached.SrcIntePseRequest) {
	fmt.Println("this is uipInterFunc pse method")
	var keysGet Pse
	var responseStr []byte
	var resopnseData responseJsonForUipInterFuncsGet
	requestStr := " {\"com\":\"search\",\"data\":{\"tableName\":\"uipInterFunc"+data.GroupId+"\",\"keyWords\":\""+data.Settext+"\"}}"
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
			if err != nil || res.GroupId != data.GroupId{
				fmt.Println("in pse method, some err occured:",err)
				continue
			}
			resopnseData.Data = append(resopnseData.Data, res)
		}
	}
	resopnseData.Code = common.ErrorOKId
	resopnseData.Msg = common.ErrorOKGetMsg
	responseStr, _ = json.Marshal(resopnseData)
	w.Write([]byte(responseStr))
}

/*
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipInterFuncvBatchQueryRequest(w http.ResponseWriter, r *http.Request, datas attached.BatchResQuest) {
	fmt.Println("this is uipInterFunc batchQuery method")
	var responseStr []byte
	var interFuncResponse responseJsonForUipInterFuncsGet
	var querySucc bool
	groupId := datas.GroupId
	for _,data := range  datas.Data{
		result, err :=frame.DB.RetriveOne("uipInterFunc"+groupId+data)
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
			querySucc = true
			interFuncResponse.Data = append(interFuncResponse.Data,uip)
		}
	}
	if querySucc {
		interFuncResponse.Code = common.ErrorOKId
		interFuncResponse.Msg = common.ErrorOKGetMsg
		responseStr, _ = json.Marshal(interFuncResponse)
		w.Write([]byte(responseStr))
	}else {
		interFuncResponse.Code = common.ErrorQueryId
		interFuncResponse.Msg = common.ErrorQueryMsg
		responseStr, _ = json.Marshal(interFuncResponse)
		w.Write([]byte(responseStr))
	}
}