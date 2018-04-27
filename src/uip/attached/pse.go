package attached

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"errors"
)

//用于接受pse返回的信息
type requestPseValue struct {
	Code string
	Msg  string
	Data []string
}

func GetPse(tableName, groupId, findValue string) ([]string, error) {
	fmt.Println("***GetPse***")

	//用于接受查询出来的key数组
	var keysGet requestPseValue

	//发送的数据
	sendValue := " {\"com\":\"search\",\"data\":{\"tableName\":\"" + tableName + groupId + "+\",\"keyWords\":\"" + findValue + "\"}}"

	//请求发送
	resp, err := http.Post("http://172.16.0.14:31007/pse",
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(sendValue)))
	//关闭
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.New("pse请求失败:" + err.Error())
	}

	//接受值并序列化
	if err := json.NewDecoder(resp.Body).Decode(&keysGet); err != nil {
		return nil, errors.New("反序列化错误:" + err.Error())
	}

	//如果数据我空则返回
	if keysGet.Data == nil {
		return nil, nil
	}

	//返回data中的key数组
	return keysGet.Data, nil
}
