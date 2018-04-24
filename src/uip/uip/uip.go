package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

/** 姑且作为示例方法，开发完成删除
*
 */
func misposHandler(w http.ResponseWriter, r *http.Request) {
	Logger.Info("in misposHandler")
	setHeader(w)

	if r.Method == "POST" {
		var request server.MisPosInfoRequestJson
		var response server.MisPosInfoRespJson
		var responseStr []byte
		requestStr := GetDataString(r)
		Logger.Info("requestStr:%s", requestStr)
		err := json.Unmarshal([]byte(requestStr), &request)

		if err != nil {
			response.Code = "10001"
			response.Msg = "json格式错误"
			Logger.Error("err:%v", err.Error())
			responseStr, _ = json.Marshal(response)
			w.Write([]byte(responseStr))
			return
		}

		command := request.Com
		data := request.Data
		switch {
		case command == "POST":
			//TODO: 调用增加逻辑
			//      请求参数有： data,是一个字符串形式的json格式数据，内容是多条完整的主表记录。
			//     修改操作需要记录操作日志
			Logger.Info("data:%v", data)
			server.MisPosPostRequest(w, data)
		case command == "DELETE":
			//TODO: 根据字典编码删除对应的数据并且 保存对应的操作日志
			//      请求参数有: groupId,是字典的机构标识数据，dictCode,是字典编码
			//   bHandleDelRequest(w, r, data)
		case command == "PUT":
			//TODO: 调用修改逻辑
			//      请求参数有：data,是一个字符串形式的json格式数据，内容是一条完整的主表记录。
			//     修改操作需要记录操作日志
			//    bHandlePutRequest(w, r, data)

			//  case command == "GETALL":
			// mALLHandleGetRequest(w, r)
		}
	} else if r.Method == "GET" {
		r.ParseForm()
		com := r.Form["com"][0]
		Logger.Info("----%v----", com)
		switch {
		case com == "ID":
			//TODO: 查询指定模块
			//	bIDHandleGetRequest(w, r)
		case com == "ALL":
			//TODO: 查询所有模块
			//     bALLHandleGetRequest(w, r)
		}
	} else {
		Logger.Info("No method or Command!")

	}
}

/**
* Default Request Handler
 */
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in defaultHandler()")
	w.Write([]byte("in defaultHandler()"))
}

/**
获取body的data(json)转换为string
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
/*
	请求头条
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
	http.HandleFunc("/", defaultHandler)
	//示例
	http.HandleFunc("/mispos", misposHandler)

}
