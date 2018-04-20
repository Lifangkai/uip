package main

import (
//	"os"
	"net/http"
	"strconv"
	"github.com/seefan/gossdb"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"server"
	."components"
	"configmsg"
	"github.com/seefan/gossdb/conf"
	"fmt"
	/*"crypto/x509"
	"crypto/tls"*/
)



func setHeader(w http.ResponseWriter){
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

/** misposHandler()
* operation of base info table
 */
func misposHandler(w http.ResponseWriter, r *http.Request) {
	Logger.Info("in misposHandler" )
	setHeader(w)

	if r.Method == "POST" {
		var request server.MisPosInfoRequestJson
		var response  server.MisPosInfoRespJson
		var responseStr []byte
		requestStr := GetDataString(r)
		Logger.Info("requestStr:%s",requestStr )
		err := json.Unmarshal([]byte(requestStr), &request)

		if err != nil {
			response.Code = "10001"
			response.Msg = "json格式错误"
			Logger.Error("err:%v", err.Error() )
			responseStr, _ = json.Marshal(response)
			w.Write([]byte(responseStr))
			return
		}

		command := request.Com
		data := request.Data
		switch{
		case command == "POST":
			//TODO: 调用增加逻辑
			//      请求参数有： data,是一个字符串形式的json格式数据，内容是多条完整的主表记录。
			//     修改操作需要记录操作日志
			Logger.Info("data:%v", data )
			server.MisPosPostRequest(w, data)
		case command == "DELETE":
			//TODO: 根据字典编码删除对应的数据并且 保存对应的操作日志
			//      请求参数有: groupId,是字典的机构标识数据，dictCode,是字典编码
			//   bHandleDelRequest(w, r, data)
		case command == "PUT" :
			//TODO: 调用修改逻辑
			//      请求参数有：data,是一个字符串形式的json格式数据，内容是一条完整的主表记录。
			//     修改操作需要记录操作日志
			//    bHandlePutRequest(w, r, data)

			//  case command == "GETALL":
			// mALLHandleGetRequest(w, r)
		}
	}else if r.Method == "GET" {
		r.ParseForm()
		com := r.Form["com"][0]
		Logger.Info("----%v----" ,com)
		switch{
		case com == "ID":
			//TODO: 查询指定模块
			//	bIDHandleGetRequest(w, r)
		case com == "ALL":
			//TODO: 查询所有模块
			//     bALLHandleGetRequest(w, r)
		}
	}else{
		Logger.Info("No method or Command!")

	}
}

/**
* Default Request Handler
 */
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	Logger.Info("in defaultHandler()")
}

/**
 * 命令行输入参数提示
 */
var Alert = func() {
	Logger.Info("Please input dbIP,dbPort,svcPort" )
}

/**
获取body的data(json)转换为string
*字节数据转string
*/
func GetDataString(req *http.Request) string {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "{\"code\": 10000,\"msg\": \"failed\"}"
	}else{
		return bytes.NewBuffer(result).String()
	}
}

/**main()
*  程序运行的入口
*  命令行参数说明：
*    dbIP: 数据库的IP 地址
*    dbPort: 数据库的端口
*    svcPort: 对外提供rest服务的端口
*/
func main() {
	//TODO : ip 和 port改成命令行输入  dbIP   dbPort
	/*args := os.Args
	if args == nil || len(args) <3{
		Alert()//如果用户没有输入,或参数个数不够,则调用该函数提示用户
		//return
	}
	//	dbIP, port, svcPort:= args[1],args[2],args[3]
	dbIP, port, svcPort:= "172.16.0.14","6100","8080"*/
    var dbIP,port,svcPort string
	//获取配置文件信息
	configmsg.InitConfig()
	Logger.Info("配置文件信息:[%v]",configmsg.Config.Udipt)
	for _, v := range configmsg.Config.Udipt {
		dbIP = v.DbIP
		port = v.Port
		svcPort = v.SvcPort
	}
	Logger.Info("dbIP = %s, dbPort = %s, svcPort=%s", dbIP, port, svcPort )
	//Logger.Info("province = %s, periCode = %s, groupid=%s", province, periCode, groupid )
	dbPort, err := strconv.Atoi(port)
	if err != nil{
		Logger.Error("字符串转换成整数失败")
	}
	//fmt.Println("dbPort :",dbPort)
	/*var err1 error
	server.DB, err1 = ssdb.Connect(dbIP, dbPort)
	Logger.Info("db= [%v]", server.DB )

	if err1 != nil {
		Logger.Error("ssdb.Connect() err = [%v]\n", err)
		os.Exit(1)
	}

	defer server.DB.Close()*/
	dbpool, err := gossdb.NewPool(&conf.Config{
		Host:             dbIP,
		Port:             dbPort,
		MinPoolSize:      5,
		MaxPoolSize:      50,
		AcquireIncrement: 5,
	})
	if err != nil {
		fmt.Printf("gossdb.NewPool error:", err.Error())
		return
	}
	defer dbpool.Close()
	server.DB, err = dbpool.NewClient()
	if err != nil {
		fmt.Printf("gossdb.NewClient error:", err.Error())
		return
	}
	defer server.DB.Close()

	//web
	//static route tmpl 是在执行程序的目录下面
	http.Handle("/css/",    http.FileServer(http.Dir("./webRoot")))
	http.Handle("/images/", http.FileServer(http.Dir("./webRoot")))
	http.Handle("/js/",     http.FileServer(http.Dir("./webRoot")))

	//dynamic route
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/mispos", misposHandler)

	//双向认证
	/*pool := x509.NewCertPool()
	caCertPath := "D:/GoProjects/newhttps/src/cert/ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	s := &http.Server{
		Addr:    ":8071",
		//Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	err = s.ListenAndServeTLS("D:/GoProjects/newhttps/src/cert/server.crt", "D:/GoProjects/newhttps/src/cert/server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}*/

	//Listen   TODO: 把 10000监听端口 改成命令行输入 svcPort
	err = http.ListenAndServe(":" + svcPort, nil)
	if err != nil {
		Logger.Error("http.ListenAndServe error = [ %s ]\n", err.Error())
	}
	Logger.Info("in main(), after http.ListenAndServe().\n" )

	return
}


