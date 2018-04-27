package main

import (
	"uip/common"
	"uip/attached"
)
func UipSrcInterfaceParaFilter(data attached.UipSrcInterface) (emptyCheckResponse responseJson, isEmpty bool) {
	isEmpty = true
	if len(data.GroupId) == 0 || len(data.InteCode)==0{
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
	} else if len(data.GroupId) != 8 {
			emptyCheckResponse.Code = common.ErrorGroupIdLenId
			emptyCheckResponse.Msg = common.ErrorGroupIdLenMsg
			isEmpty = false
	}
	if len(data.InteCode) > 32 {
		emptyCheckResponse.Code = common.ErrorInterLenId
		emptyCheckResponse.Msg = common.ErrorInterLenMsg
		isEmpty = false
	}

	return emptyCheckResponse, isEmpty
}


func UipInterFuncParaFilter(data attached.UipInterFunc) (emptyCheckResponse responseJson, isEmpty bool) {
	isEmpty = true
	if len(data.GroupId) == 0 || len(data.FuncCode)==0{
		emptyCheckResponse.Code = common.ErrorDataNullId
		emptyCheckResponse.Msg = common.ErrorDataNullMsg
		isEmpty = false
	} else if len(data.GroupId) != 8 {
		emptyCheckResponse.Code = common.ErrorGroupIdLenId
		emptyCheckResponse.Msg = common.ErrorGroupIdLenMsg
		isEmpty = false
	}
	if len(data.FuncCode) > 32 {
		emptyCheckResponse.Code = common.ErrorFuncLenId
		emptyCheckResponse.Msg = common.ErrorFuncLenMsg
		isEmpty = false
	}

	return emptyCheckResponse, isEmpty
}