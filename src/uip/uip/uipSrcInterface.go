package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"uip/src/uip/common"
	// "rsf/ysrd/impl/etcd"
	// "io/ioutil"
	"io/ioutil"
	"bytes"
)


type  UipSrcInterface struct {
	Cmt	string	`json:"cmt"`
	Ctime	string	`json:"ctime"`
	UserCode	string	`json:"userCode"`
	FuncList	string	`json:"funcList"`
	OpenList	string	`json:"openList"`
	ConnConf	string	`json:"connConf"`
	InteType	string	`json:"inteType"`
	InteCode	string	`json:"inteCode"`
	InteName	string	`json:"inteName"`
	GroupId	string	`json:"groupId"`
}

type responseJson struct {
	 Code		string `json:"code"`
	 Msg		string `json:"msg"`
}
type responseJsonForUipSrcInterfaceGet struct {
	 Code		string `json:"code"`
	 Msg		string `json:"msg"`
	 Data		UipSrcInterface `json:"data"`
}
type responseJsonForUipSrcInterfacesGet struct {
	Code		string `json:"code"`
	Msg		string `json:"msg"`
	Data		[]UipSrcInterface `json:"data"`
}
type srcIntePseRequest struct {
	 GroupId		string `json:"groupId"`
	 Fttext			string `json:"fttext"`
}

type Pse struct {
	Com string `json:"com"`
	Msg string `json:"msg"`
	Data []string `json:"data"`
}
/*@Author:zengzeen
* @param:msrCountList struct
* @return:{code,msg,data}   json
 */


func UipSrcInterfaceHandlePostRequest(w http.ResponseWriter, r *http.Request, data UipSrcInterface) {
	fmt.Println("this is uipSrcInterfaceHandlePostRequest!")
	var responseInfo responseJson
	var responseStr []byte

	//获取输入数据
	dataStr, _ := json.Marshal(data)
	fmt.Println("dataStr=[", dataStr, "]")
	var uip UipSrcInterface
	err := json.Unmarshal([]byte(dataStr), &uip)

	if err != nil {
		responseInfo.Code = common.ErrorJsonFmtErrId
		responseInfo.Msg = common.ErrorJsonFmtErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	data.InteCode ,err = getSnsNew("inteCode")
	emptyCheckResponse, isEmpty := UipSrcInterfaceParaFilter(uip)
	//	_, isEmpty := msrCountListParaFilter(msr)
	if isEmpty == false {
		responseStr, _ = json.Marshal(emptyCheckResponse)
		w.Write([]byte(responseStr))
		return
	}

	//判断是否已经存在数据
	isExists, errI := aRsf.DB.ExistsDo("exists", "uipSrcInterface"+data.GroupId+data.InteCode)
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
	dbinfo, err :=aRsf.DB.Create("uipSrcInterface","uipSrcInterface"+data.GroupId+data.InteCode, string(dataStr))
	//异常处理
	if err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		fmt.Println(dbinfo)
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


/*
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipSrcInterfaceDelRequest(w http.ResponseWriter, r *http.Request, data UipSrcInterface) {
	fmt.Println("this is UipSrcInterface Del method")
	var responseInfo responseJson
	var responseStr []byte
	dataStr, _ := json.Marshal(data)
	var uip UipSrcInterface
	err := json.Unmarshal([]byte(dataStr), &uip)
	if err != nil {
		responseInfo.Code = common.ErrorJsonFmtErrId
		responseInfo.Msg = common.ErrorJsonFmtErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}

	emptyCheckResponse,isEmpty := UipSrcInterfaceParaFilter(data)
	if isEmpty == false {
		responseStr, _ = json.Marshal(emptyCheckResponse)
		w.Write([]byte(responseStr))
		return

	}
	//删除前先保存原始数据
	OldData, err := aRsf.DB.RetriveOne("msrCountList" + data.GroupId + data.InteCode)
	fmt.Println("OldData = [", OldData, "]")
	if err == nil && OldData == nil {
		responseInfo.Code = common.ErrorDataNotExistsErrId
		responseInfo.Msg = common.ErrorDataNotExistsMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
	} else if OldData == nil && err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))

	} else {
		dbinfo, err :=aRsf.DB.Delete("uipSrcInterface","uipSrcInterface" + data.GroupId+data.InteCode)
		//异常处理
		if err != nil {
			responseInfo.Code = common.ErrorSystemErrId
			responseInfo.Msg = common.ErrorSystemErrMsg
			fmt.Println(dbinfo)
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

func UipSrcInterfaceHandlePutRequest(w http.ResponseWriter, r *http.Request, data UipSrcInterface) {
	fmt.Println("this is UipSrcInterface Put method")
	var responseInfo responseJson
	var responseStr []byte

	dataStr, _ := json.Marshal(data)
	var uip UipSrcInterface
	err := json.Unmarshal([]byte(dataStr), &uip)
	if err != nil {
		responseInfo.Code = common.ErrorJsonFmtErrId
		responseInfo.Msg = common.ErrorJsonFmtErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	emptyCheckResponse, isEmpty := UipSrcInterfaceParaFilter(uip)
	if isEmpty == false {
		responseStr, _ = json.Marshal(emptyCheckResponse)
		w.Write([]byte(responseStr))
		return
	}
	//判断是否已经存在数据
	OldData, err := aRsf.DB.RetriveOne("uipSrcInterface" + data.GroupId + data.InteCode)
	if err == nil && OldData == nil {
		responseInfo.Code = common.ErrorDataNotExistsErrId
		responseInfo.Msg = common.ErrorDataNotExistsMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
	if OldData == nil && err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	} else {
		fmt.Println("UipSrcInterface = [", uip, "]")
		//向数据库插入新数据
		dbinfo, err := aRsf.DB.Update("uipSrcInterface","uipSrcInterface"+data.GroupId+data.InteCode, string(dataStr))
		//异常处理
		if err != nil {
			responseInfo.Code = common.ErrorSystemErrId
			responseInfo.Msg = common.ErrorSystemErrMsg
			fmt.Println(dbinfo)
			responseStr, _ = json.Marshal(responseInfo)
			w.Write([]byte(responseStr))
		}
		responseInfo.Code = common.ErrorOKId
		responseInfo.Msg = common.ErrorOKModifyMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	}
}

/** GET方法 com=SN 根据groupId和sn查询数据
*	@Author zengzeen
*	@param  groupId 组织标识
*	@param  key 前来传过来用于查询的sn，即序列号
*   @return {code,msg,data} json
 */
func UipSrcInterfaceGetRequest(w http.ResponseWriter, r *http.Request ) {
	fmt.Println("this is UipSrcInterface SNGet method")
	var responseInfo responseJson
	var responseStr []byte
	var responseInfoForUipSrcInterface responseJsonForUipSrcInterfaceGet
	var groupId,inteCode string

	groupId = r.Form["groupId"][0]
	inteCode = r.Form["inteCode"][0]

	if len(groupId)==0 || len(inteCode)==0 {
		responseInfo.Code = common.ErrorDataNullId
		responseInfo.Msg = common.ErrorDataNullMsg
		return
	}else if len(groupId)!=8 {
		responseInfo.Code = common.ErrorGroupIdLenId
		responseInfo.Msg = common.ErrorGroupIdLenMsg
		return
	}else if len(inteCode)>32 {
		responseInfo.Code = common.ErrorInterLenId
		responseInfo.Msg = common.ErrorInterLenMsg
		return
	}

	//根据主键搜索
	result, err :=aRsf.DB.RetriveOne("msrCountList"+groupId+inteCode)
	//fmt.Println("result.(type): ", reflect.TypeOf(result))
	if err == nil && result == nil {
		responseInfo.Code = common.ErrorDataNotExistsErrId
		responseInfo.Msg = common.ErrorDataNotExistsMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	} else if result == nil && err != nil {
		responseInfo.Code = common.ErrorSystemErrId
		responseInfo.Msg = common.ErrorSystemErrMsg
		responseStr, _ = json.Marshal(responseInfo)
		w.Write([]byte(responseStr))
		return
	} else {
		responseInfo.Code = common.ErrorOKId
		responseInfo.Msg = common.ErrorOKSelectMsg
		var uip UipSrcInterface

		err := json.Unmarshal([]byte(result.(string)), &uip)
		if err != nil {
			var response responseJson
			response.Code = common.ErrorJsonFmtErrId
			response.Msg = common.ErrorJsonFmtErrMsg
			responseStr, _ = json.Marshal(response)
			w.Write([]byte(responseStr))
			return
		}
		responseInfoForUipSrcInterface.Data = uip
		responseInfoForUipSrcInterface.Code = common.ErrorOKId
		responseInfoForUipSrcInterface.Msg = common.ErrorOKGetMsg
		responseStr, _ = json.Marshal(responseInfoForUipSrcInterface)
		w.Write([]byte(responseStr))
	}
}

/**
*	@Author:zengzeen
*	@param:groupId          机构标识     定长八位
*   @param:Sn               序号
*   @return:{code,msg,data} json
 */

func UipSrcInterfacePseHandlePutRequest(w http.ResponseWriter, r *http.Request, data srcIntePseRequest) {
	fmt.Println("this is UipSrcInterface pse method")
	var responseInfo responseJson
	var responseStr []byte
	var keysGet Pse
	requestStr := " {\"com\":\"search\",\"data\":{\"tableName\":\"uipSrcInterface"+data.GroupId+"\",\"keyWords\":\""+data.Fttext+"\"}}"
	var valstr=[]byte(requestStr)
	resp, err := http.Post("http://172.16.0.14:31007/pse",
		"application/json;charset=utf-8", bytes.NewBuffer(valstr))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	data2,err2:= ioutil.ReadAll(resp.Body)
	fmt.Println("the err when we read the respons's body is :",err2)
	keys := keysGet.Data
	for _, keyword := range keys {
		var res UipSrcInterface
		childRes := getlistOne(keyword)
		fmt.Println(childRes)
		if len(childRes.Data) != 0 {
			err := json.Unmarshal([]byte(childRes.Data), &res)
			//异常处理
			if err != nil || res.GroupId != groupID{
				fmt.Println("inChildTable's pseFunc methods","the err is :",err)
				continue
			}
			subs = append(subs, res)
		}
	}
}
/** GET方法 com=POST 根据关键字查询
*	@Author zengzeen
*	@param  groupId 组织标识
*	@param  fttext 搜索关键字
*   @return {code,msg,data} json
 */

func PseSerach(keyword string,dictCode string, groupID string) (uips []UipSrcInterface) {
	requestStr := " {\"com\":\"search\",\"data\":{\"tableName\":\"dictClassSub"+dictCode+"\",\"keyWords\":\""+keyword+"\"}}"
	fmt.Println("the requestStr is :",requestStr)
	var valstr=[]byte(requestStr)
	var keysGet Pse
	resp, err := http.Post("http://172.16.0.14:31007/pse",
		"application/json;charset=utf-8", bytes.NewBuffer(valstr))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	data,err3:= ioutil.ReadAll(resp.Body)
	fmt.Println("the data we got after post is:",string(data))
	fmt.Println("the err when post for pse is :",err3)
	err =json.Unmarshal([]byte(data),&keysGet)
	resp.Body.Close()
	if err !=nil{
		fmt.Println("the json umarshal to keys[] is faild, the err is:",err)
	}
	keys := keysGet.Data
	for _, keyword := range keys {
		var res DictClassSub
		childRes := getlistOne(keyword)
		fmt.Println(childRes)
		if len(childRes.Data) != 0 {
			err := json.Unmarshal([]byte(childRes.Data), &res)
			//异常处理
			if err != nil || res.GroupId != groupID{
				fmt.Println("inChildTable's pseFunc methods","the err is :",err)
				continue
			}
			subs = append(subs, res)
		}
	}
	return subs
}