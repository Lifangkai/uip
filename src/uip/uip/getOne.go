package main

import (
	"uip/common"
	"fmt"
)

/**responseJson（）
*	返回给前端的json数据
 */
type responJson struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

/*@Author:zhengxianbo
* @param:keyword 			  传入的字符串
* @return:responsesChild      responseJson
 */
//getlistOne()函数通过传入key来查询子表信息
func getlistOne(keyword string) (responsesChild responJson) {

	fmt.Println("this is childTable getlistOne method")
	key := keyword
	fmt.Println("key = [", key, "]")
	//根据key查询所有数据dictCode + itemCode搜索
	result, err := frame.DB.RetriveOne(key)

	fmt.Println("result = [", result, "]")
	//异常处理
	if err == nil && result == "" {
		responsesChild.Code = common.ErrorDataNotExistsErrId
		responsesChild.Msg = common.ErrorDataNotExistsMsg
		return responsesChild
	} else if result == "" && err != nil {
		responsesChild.Code = common.ErrorSystemErrId
		responsesChild.Msg = common.ErrorSystemErrMsg
		return responsesChild
	} else {
		responsesChild.Data = result
		return responsesChild
	}
}
