package main

import (
	"uip/common"
	"uip/attached"
)

/**UipSrcInterfaceParaFilter()
*
*
*/
func UipSrcInterfaceParaFilter(data attached.UipSrcInterface) (emptyCheckResponse responseJson, isEmpty bool) {
	common.Frame.Log.Write("in UipSrcInterfaceParaFilter()")

	isEmpty = true

	if len(data.GroupId) == 0 || len(data.InteCode)==0 {
		common.Frame.Log.Write(common.ErrorDataNullMsg)
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.GroupId) != 8 {
		common.Frame.Log.Write(common.ErrorGroupIdLenMsg)
		emptyCheckResponse.Code = common.ErrorGroupIdLenId
		emptyCheckResponse.Msg = common.ErrorGroupIdLenMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.InteCode) > 32 {
		common.Frame.Log.Write(common.ErrorInterLenMsg)
		emptyCheckResponse.Code = common.ErrorInterLenId
		emptyCheckResponse.Msg = common.ErrorInterLenMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	return emptyCheckResponse, isEmpty
}

/**UipInterFuncParaFilter（）
* 检查功能配置的输入参数的有效性
*/
func UipInterFuncParaFilter(data attached.UipInterFunc) (emptyCheckResponse responseJson, isEmpty bool) {
	isEmpty = true

	if len(data.GroupId) == 0 {
		common.Frame.Log.Write("groupId = 0")
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.FuncCode)==0 {
		common.Frame.Log.Write("funcode = 0")
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.InteCode) == 0 {
		common.Frame.Log.Write("inteCode = 0")
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.ResFmtCode) == 0 {
		common.Frame.Log.Write("resFmtCode = 0")
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.ReqFmtCode) == 0{
		common.Frame.Log.Write("reqFmtCode = 0")
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.GroupId) != 8 {
		common.Frame.Log.Write("groupId!= 8")
		emptyCheckResponse.Code = common.ErrorGroupIdLenId
		emptyCheckResponse.Msg = common.ErrorGroupIdLenMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	if len(data.FuncCode) > 32 {
		common.Frame.Log.Write("FuncCode > 32")
		emptyCheckResponse.Code = common.ErrorFuncLenId
		emptyCheckResponse.Msg = common.ErrorFuncLenMsg
		isEmpty = false
		return emptyCheckResponse, isEmpty
	}

	return emptyCheckResponse, isEmpty
}

