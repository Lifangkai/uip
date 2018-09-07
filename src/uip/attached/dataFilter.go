package attached

import (
	"fmt"
  "uip/common"
)

/*
	主表数据过滤方法
*/
func MainFilter(uipMain UipDfmtMain) bool {
  common.Frame.Log.Write("in MainFilter(): uipMain = " + fmt.Sprint(uipMain))
	var boo bool = true

	//为空判断
	if uipMain.GroupId == "" || uipMain.FmtType == ""  || uipMain.OperCode == "" {
		boo = false
		common.Frame.Log.Write("有字段为空！")
		return boo
	}

	//长度判断
	if len(uipMain.GroupId) != 8 || len(uipMain.FmtName) > 64 || len(uipMain.FmtType) > 16 || len(uipMain.ExtInfo) > 256 || len(uipMain.OperCode) > 64 || len(uipMain.Cmt) > 256 {
		boo = false
		common.Frame.Log.Write("长度不符合！")
		return boo
	}

	//默认值判断
	if uipMain.FmtType != "json-json" && uipMain.FmtType != "file-file" && uipMain.FmtType != "ws-web" && uipMain.FmtType != "service" {
		boo = false
		common.Frame.Log.Write("默认值不合法！")
		return boo
	}

	return boo
}

/**
*	数据格式子表数据过滤方法
*/
func SubFilter(uipSub UipDfmtSub) bool {
  common.Frame.Log.Write("in SubFilter() uipSub = " + fmt.Sprint(uipSub))
	var boo bool = true

	//为空判断
	if uipSub.GroupId == ""{ 
    common.Frame.Log.Write("groupId is null")
		boo = false
		return boo
	}

	if uipSub.FmtCode == ""{
    common.Frame.Log.Write("FmtCode is null")
		boo = false
		return boo
	}

	if uipSub.FieldName == ""{
    common.Frame.Log.Write("Field Name is null")
		boo = false
		return boo
	}

	if uipSub.DataType == "" {
    common.Frame.Log.Write("DataType is null")
		boo = false
		return boo
	}

	if uipSub.Length == "" {
    common.Frame.Log.Write("Length is null")
		boo = false
		return boo
	}

	if uipSub.Constrain == "" {
    common.Frame.Log.Write("Constrain is null")
		boo = false
		return boo
	}

	if uipSub.ExtInfo == ""{
    common.Frame.Log.Write("ExtInfo is null")
		boo = false
		return boo
	}

	if uipSub.OperCode == "" {
    common.Frame.Log.Write("OperCode is null")
		boo = false
		return boo
	}

	//长度判断
	if len(uipSub.GroupId) != 8 {
		boo = false
		common.Frame.Log.Write("groupId 长度不等于8")
		return boo
	}

	if len(uipSub.FmtCode) > 32 {
		boo = false
		common.Frame.Log.Write("FmtCode 长度大于32")
		return boo
	}

	if len(uipSub.DtlCode) > 32 {
		boo = false
		common.Frame.Log.Write("DtlCode 长度大于32")
		return boo
	}

	if len(uipSub.FieldName) > 50 {
		boo = false
		common.Frame.Log.Write("FieldName 长度大于50")
		return boo
	}

	if len(uipSub.DataType) > 16 {
		boo = false
		common.Frame.Log.Write("DataType 长度大于16")
		return boo
	}

	if len(uipSub.Length) > 16 {
		boo = false
		common.Frame.Log.Write("Length 长度大于16")
		return boo
	}

	if len(uipSub.Constrain) > 16 {
		boo = false
		common.Frame.Log.Write("Constrain 长度大于16")
		return boo
	}

	if len(uipSub.ExtInfo) > 256 {
		boo = false
		common.Frame.Log.Write("ExtInfo 长度大于256")
		return boo
	}

	if len(uipSub.OperCode) > 64 {
		boo = false
		common.Frame.Log.Write("OperCode 长度大于64")
		return boo
	}

	if len(uipSub.Cmt) > 256 {
		boo = false
		common.Frame.Log.Write("Cmt 长度大于256")
		return boo
	}

	return boo
}

/**
*	接口开放过滤方法
*/
func OpenInterfaceFilter(data UipOpenInterface) bool {
	common.Frame.Log.Write("in OpenInterfaceFilter(): data = " + fmt.Sprint(data))

	var boo bool = true

	//为空判断
	if data.GroupId == "" || data.InteCode == "" || data.InteType == "" || data.ConnConf == "" || data.SinteCode == "" || data.OperCode == "" {
		boo = false
		fmt.Println("有字段为空！")
		return boo
	}

	//长度判断
	if len(data.GroupId) != 8 || len(data.InteName) > 64 || len(data.InteCode) > 16 || len(data.ConnConf) > 14 || len(data.FuncList) > 1024 || len(data.BeforeCode) > 32 || len(data.SinteCode) > 32 || len(data.AfterCode) > 32 || len(data.OperCode) > 64 || len(data.Cmt) > 256 {
		boo = false
		fmt.Println("长度不符合！")
		return boo
	}

	//默认值判断
	if data.InteType != "Oracle" && data.InteType != "http" && data.InteType != "SQLServer" {
		boo = false
		fmt.Println("默认值不合法！")
		return boo
	}

	return boo
}
