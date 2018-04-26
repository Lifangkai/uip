package attached

import (
	"fmt"
)

/*
	主表数据过滤方法
*/
func MainFilter(uipMain UipDfmtMain) bool {
	var boo bool = true

	//为空判断
	if uipMain.GroupId == "" || uipMain.FmtType == "" || uipMain.ExtInfo == "" || uipMain.OperCode == "" {
		boo = false
		fmt.Println("有字段为空！")
		//res:=&response{
		//	Code:common.CheckErrorId,
		//	Msg:common.CheckErrorMsg,
		//}
		//res.Answer(w)
		return boo
	}

	//长度判断
	if len(uipMain.GroupId) != 8 || len(uipMain.FmtName) > 64 || len(uipMain.FmtType) > 16 || len(uipMain.ExtInfo) > 256 || len(uipMain.OperCode) > 64 || len(uipMain.Cmt) > 256 {
		boo = false
		fmt.Println("长度不符合！")
		return boo
	}

	//默认值判断
	if uipMain.FmtType != "json-json" && uipMain.FmtType != "file-file" && uipMain.FmtType != "ws-web" && uipMain.FmtType != "service" {
		boo = false
		fmt.Println("默认值不合法！")
		return boo
	}

	return boo
}

/*
	子表数据过滤方法
*/
func SubFilter(uipSub UipDfmtSub) bool {
	var boo bool = true

	//为空判断
	if uipSub.GroupId == "" || uipSub.FmtCode == "" || uipSub.FieldName == "" || uipSub.DataType == "" || uipSub.Length == "" || uipSub.Constrain == "" || uipSub.ExtInfo == "" || uipSub.OperCode == "" {
		boo = false
		fmt.Println("有字段为空！")
		//res:=&response{
		//	Code:common.CheckErrorId,
		//	Msg:common.CheckErrorMsg,
		//}
		//res.Answer(w)
		return boo
	}

	//长度判断
	if len(uipSub.GroupId) != 8 || len(uipSub.FmtCode) > 32 || len(uipSub.DtlCode) > 32 || len(uipSub.FieldName) > 50 || len(uipSub.DataType) > 16 || len(uipSub.Length) > 16 || len(uipSub.Constrain) > 16 || len(uipSub.ExtInfo) > 256 || len(uipSub.OperCode) > 64 || len(uipSub.Cmt) > 256 {
		boo = false
		fmt.Println("长度不符合！")
		return boo
	}
	return boo
}
