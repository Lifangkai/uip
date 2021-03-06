package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"

	"rsf/rsf"
	"uip/common"
	"uip/attached"
)

var (
	//定义全局变量指针Rsf
	//Frame *rsf.Rsf
	//定义响应结构体
	response *attached.Response
)

/*
* describe:接口数据格式信息总表/主表
*/
func dfmtMainHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in dfmtMainHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.MainRequestJson

		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}

		//根据传入的对应状态值执行相应的方法
		switch request.Com {
		case "POST":
			mainInsertValueMethod(w, request.Data)
		case "PUT":
			mainUpdateValueMethod(w, request.Data)
		case "DELETE":
			mainDeleteValueMethod(w, request.Data)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}

	//GET逻辑
	if r.Method == "GET" {
		//解析
		if err := r.ParseForm(); err != nil {
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}

		//判断参数是否正确
		comWay := r.Form["com"]
		if len(comWay) != 1 || comWay[0] == "" {
			response.Code = common.ErrorParameterIsErrId
			response.Msg = common.ErrorParameterIsErrMsg
		}

		//根据传入的参数执行对应的方法
		switch comWay[0] {
		case "search":
			mainQueryValueMethod(w, r)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/*
* describe:接口数据格式信息子表
*/
func dfmtSubHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in dfmtSubHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.SubRequestJson

		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}

		//根据传入的对应状态值执行相应的方法
		switch request.Com {
		case "POST":
			subInsertValueMethod(w, request.Data)
		case "PUT":
			subUpdateValueMethod(w, request.Data)
		case "DELETE":
			subDeleteValueMethod(w, request.Data)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/*
* describe:接口开放管理
*/
func uipOpenInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipOpenInterfaceHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.RequestOpenInterfaceJson

		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}

		//根据传入的对应状态值执行相应的方法
		switch request.Com {
		case "POST":
			interfaceOpenInsertMethod(request.Data)
		case "PUT":
			interfaceOpenUpdateMethod(request.Data)
		case "DELETE":
			interfaceOpenDeleteMethod(request.Data)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
	
	//GET逻辑
	if r.Method == "GET" {
		//解析
		if err := r.ParseForm(); err != nil {
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}

		//判断参数是否正确
		comWay := r.Form["com"]
		if len(comWay) != 1 || comWay[0] == "" {
			response.Code = common.ErrorParameterIsErrId
			response.Msg = common.ErrorParameterIsErrMsg
		}

		//根据传入的参数执行对应的方法
		switch comWay[0] {
		case "one":
			interfaceOpenQueryOneMethod(r)
		case "list":
			//全文检索查询出前20条数据
			interfaceOpenQueryManyMethod(r)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/**
* describe:源接口管理表格基本方法
*/
func uipSrcInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipSrcInterfaceHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.UipSrcInteRequestJson

		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}

		//根据传入的对应状态值执行相应的方法
		switch request.Com {
		case "POST":
			UipSrcInterfaceHandlePostRequest(request.Data)
		case "PUT":
			UipSrcInterfaceHandlePutRequest(request.Data)
		case "DELETE":
			UipSrcInterfaceDelRequest(request.Data)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}

	//GET逻辑
	if r.Method == "GET" {
		//解析
		if err := r.ParseForm(); err != nil {
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}

		//判断参数是否正确
		comWay := r.Form["com"]
		if len(comWay) != 1 || comWay[0] == "" {
			response.Code = common.ErrorParameterIsErrId
			response.Msg = common.ErrorParameterIsErrMsg
		}

		//根据传入的参数执行对应的方法
		switch comWay[0] {
		case "key":
			UipSrcInterfaceGetRequest(r)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/**
* describe:源接口管理全文搜索方法
*/
func uipSrcInterfaceFtsearchHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipSrcInterfaceFtsearchHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	if r.Method == "GET" {
		//解析
		if err := r.ParseForm(); err != nil {
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}

		//判断参数是否正确
		comWay := r.Form["com"]
		if len(comWay) != 1 || comWay[0] == "" {
			response.Code = common.ErrorParameterIsErrId
			response.Msg = common.ErrorParameterIsErrMsg
		}

		//根据传入的参数执行对应的方法
		switch comWay[0] {
		case "key":
			UipSrcInterfacePseHandlePutRequest(r)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/**
* describe:接口功能表格基本方法
*/
func uipInterFuncHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipInterFuncHandler（）")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.UipInterFuncRequestJson

		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}

		//根据传入的对应状态值执行相应的方法
		switch request.Com {
		case "POST":
			UipInterFuncHandlePostRequest(request.Data)
		case "PUT":
			UipInterFuncHandlePutRequest( request.Data)
		case "DELETE":
			UipInterFuncvDelRequest(request.Data)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}

	//GET逻辑
	if r.Method == "GET" {
		//解析
		if err := r.ParseForm(); err != nil {
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}

		//判断参数是否正确
		comWay := r.Form["com"]
		if len(comWay) != 1 || comWay[0] == "" {
			response.Code = common.ErrorParameterIsErrId
			response.Msg = common.ErrorParameterIsErrMsg
		}

		//根据传入的参数执行对应的方法
		switch comWay[0] {
		case "key":
			UipInterFuncGetRequest(r)
		case "all":
			UipInterFuncQuryAlltRequest(r)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/**
* describe:接口功能管理全文搜索方法
*/
func uipInterFuncFtsearchHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipInterFuncFtsearchHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	if r.Method == "GET" {
		//解析
		if err := r.ParseForm(); err != nil {
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}

		//判断参数是否正确
		comWay := r.Form["com"]
		if len(comWay) != 1 || comWay[0] == "" {
			response.Code = common.ErrorParameterIsErrId
			response.Msg = common.ErrorParameterIsErrMsg
		}

		//根据传入的参数执行对应的方法
		switch comWay[0] {
		case "key":
			UipInterFuncPseHandlePutRequest(r)
		default:
			response.Code = common.ComErrorId
			response.Msg = common.ComErrorMsg
			return
		}
	}
}

/**
* describe:接口功能管理批量增加方法
*/
func uipInterFuncBatchCreateHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipInterFuncBatchCreateHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.UipInterFuncsRequestJson
		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}
		UipInterFuncsHandlePostRequest(request.Data)
	}
}

/**
* describe:接口功能管理批量搜索方法
*/
func uipInterFuncBSearchHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in uipInterFuncFtsearchHandler()")

	//实例化给前段的状态的对象
	response = attached.NewResponse()

	//如果没有对应的请求方式则请求错误,同时执行对应的响应信息
	defer response.Answer(w)

	//设置请求方法
	setHeader(w)

	//POST逻辑
	if r.Method == "POST" {
		//用于接受前段给的值
		var request attached.BatchResQuest
		//接受并解析数据
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Code = common.ErrorJsonErrId
			response.Msg = common.ErrorJsonErrMsg
			return
		}
		UipInterFuncvBatchQueryRequest(request)
	}
}

/**defaultHandler()
* Default Request Handler
*/
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	common.Frame.Log.Write("in defaultHandler()")
	w.Write([]byte("in defaultHandler()"))
}

/**GetDataString()
*获取body的data(json)转换为string
*字节数据转string
*/
func GetDataString(req *http.Request) string {
	defer req.Body.Close()
	
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "{\"code\": 10000,\"msg\": \"failed\"}"
	} else {
		return bytes.NewBuffer(result).String()
	}
}

/**
* 请求头
*/
func setHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

/*main()
*  程序运行的入口
*/
func main() {

	//web
	//static route tmpl 是在执行程序的目录下面
	http.Handle("/css/", http.FileServer(http.Dir("./webRoot")))
	http.Handle("/images/", http.FileServer(http.Dir("./webRoot")))
	http.Handle("/js/", http.FileServer(http.Dir("./webRoot")))

	//dynamic route
	common.Frame = rsf.NewRsf()
	
	//common.AExt = aYljg

  err := json.Unmarshal([]byte( common.Frame.Ext), &common.AExt)
  if err != nil {
    common.Frame.Log.Write("get ext parameter object err=[" + err.Error() + "], ext = [" + common.Frame.Ext +"]")
    return
    }
	
	//默认路由
	common.Frame.SetPathHandlerPair("/", defaultHandler)

	//接口数据格式信息总表/主表 uip_dfmt_main
	common.Frame.SetPathHandlerPair("/fmt/manage/main", dfmtMainHandler)

	//接口数据格式信息子表 uip_dfmt_sub
	common.Frame.SetPathHandlerPair("/fmt/manage/sub", dfmtSubHandler)

	//uipSrcInteface源接口管理
	common.Frame.SetPathHandlerPair("/sinte/manage", uipSrcInterfaceHandler)
	common.Frame.SetPathHandlerPair("/sinte/manage/ftsearch", uipSrcInterfaceFtsearchHandler)

	//uipInterFunc接口功能管理表
	common.Frame.SetPathHandlerPair("/func/manage", uipInterFuncHandler)
	common.Frame.SetPathHandlerPair("/func/manage/ftsearch", uipInterFuncFtsearchHandler)
	common.Frame.SetPathHandlerPair("/func/manage/batchCreate",uipInterFuncBatchCreateHandler)
	common.Frame.SetPathHandlerPair("/func/manage/bSearch", uipInterFuncBSearchHandler)

	//开放接口管理 uip_open_interface
	common.Frame.SetPathHandlerPair("/ointe/manage", uipOpenInterfaceHandler)

	common.Frame.Run()

}
