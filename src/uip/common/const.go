package common

const (
	//服务调用没有发生任何错误
	ErrorOKId        = "200000" //所有的服务都以这个编码为正确无误调用之后的返回值
	ErrorOKMsg       = "OK"     //所有的服务都以这个消息为200000错误码对应的消息
	ErrorOKInsertMsg = "增加成功"
	ErrorOKDeleteMsg = "删除成功"
	ErrorOKModifyMsg = "修改成功"
	ErrorOKSelectMsg = "查询成功"
	ErrorOKGetMsg    = "数据查询成功"

	ComErrorId  = "200001"
	ComErrorMsg = "com参数错误"

	InsertErrorId  = "200002"
	InsertErrorMsg = "插入失败"

	CheckErrorId  = "200003"
	CheckErrorMsg = "数据未通过校验"

	//Json格式错误
	ErrorJsonErrId  = "200004"
	ErrorJsonErrMsg = "json格式错误"

	//get输入com判断
	ErrorComisNilId  = "200005"
	ErrorComisNilMsg = "查询命令标识为空"

	//数据异常
	ErrorId  = "200006"
	ErrorMsg = "数据异常"

	//数据已存在
	ErrorDataExistsErrId  = "200007"
	ErrorDataExistsErrMsg = "数据已存在"

	//系统错误
	ErrorSystemErrId  = "200008"
	ErrorSystemErrMsg = "系统错误"

	//数据不存在
	ErrorDataNotExistsErrId = "20009"
	ErrorDataNotExistsMsg   = "数据不存在"

	ErrorDiscoverCheckId  = "200010"
	ErrorDiscoverCheckMsg = "sns服务发现失败"

	ErrorAddInfoId  = "200011"
	ErrorAddInfoMsg = "添加错误"

	ErrorUpdataInfoFailedId  = "200012"
	ErrorUpdataInfoFailedMsg = "修改失败"

	ErrorRequestBadWayId  = "200013"
	ErrorRequestBadWayMsg = "请求方式错误"

	ErrorParameterIsErrId  = "200014"
	ErrorParameterIsErrMsg = "参数错误"

	ErrorDelErrId="200015"
	ErrorDelErrMsg="删除错误"

	ErrorJsonFmtErrId = "2000018"
	ErrorJsonFmtErrMsg = "json格式错误!"

	ErrorDataNullId = "2000016"
	ErrorDataNullMsg = "请求参数为空!"

	ErrorGroupIdLenId = "2000017"
	ErrorGroupIdLenMsg ="机构表示长度不为8位！"

	ErrorInterLenId = "2000018"
	ErrorInterLenMsg = "接口编码长度大于32位！"

	ErrorFuncLenId = "2000019"
	ErrorFuncLenMsg = "接口功能编码长度大于32位！"

	ErrorQueryId = "2000020"
	ErrorQueryMsg = "一条或者多条数据查询失败！"

	ErrorQueryErrId="200021"
	ErrorQueryErrMsg="查询错误"

	ErrorInterfaceNotExsitId="200022"
	ErrorInterfaceNotExsitMsg="源接口数据不存在"

	ErrorDelSrcInterfaceErrId="200023"
	ErrorDelSrcInterfaceErrMsg="在源接口中删除开放接口编码错误:"

	ErrorSrcInterfaceDataIsNilId="200024"
	ErrorSrcInterfaceDataIsNilMsg="源接口不存在"

	ErrorOpenInterfaceDataIsNilId="200025"
	ErrorOpenInterfaceDataIsNilMsg="删除源接口表错误"
)
