package attached

//接口格式主表
type UipDfmtMain struct {
	Cmt      string `json:"cmt"`
	OperCode string `json:"operCode"`
	Ctime    string `json:"ctime"`
	//dtlCode  string `json:"dtlCode"`
	ExtInfo string `json:"extInfo"`
	FmtType string `json:"fmtType"`
	FmtCode string `json:"fmtCode"`
	FmtName string `json:"fmtName"`
	GroupId string `json:"groupId"`
}

//接口格式接受方法
type MainRequestJson struct {
	Com  string      `json:"com"`
	Data UipDfmtMain `json:"data"`
}

//接口格式子表
type UipDfmtSub struct {
	Cmt       string `json:"cmt"`
	OperCode  string `json:"operCode"`
	Ctime     string `json:"ctime"`
	ExtInfo   string `json:"extInfo"`
	Constrain string `json:"constrain"`
	Length    string `json:"length"`
	DataType  string `json:"dataType"`
	FieldName string `json:"fieldName"`
	DtlCode   string `json:"dtlCode"`
	FmtCode   string `json:"fmtCode"`
	GroupId   string `json:"groupId"`
}

//接口格式子表接受方法
type SubRequestJson struct {
	Com  string     `json:"com"`
	Data UipDfmtSub `json:"data"`
}

type UipSrcInterface struct {
	Cmt      string `json:"cmt"`
	Ctime    string `json:"ctime"`
	UserCode string `json:"userCode"`
	FuncList string `json:"funcList"`
	OpenList string `json:"openList"`
	ConnConf string `json:"connConf"`
	InteType string `json:"inteType"`
	InteCode string `json:"inteCode"`
	InteName string `json:"inteName"`
	GroupId  string `json:"groupId"`
}
type UipSrcInteRequestJson struct {
	Com  string          `json:"com"`
	Data UipSrcInterface `json:"data"`
}
type UipInterFunc struct {
	Cmt         string `json:"cmt"`
	OperCode    string `json:"operCode"`
	Ctime       string `json:"ctime"`
	ResFmtCode  string `json:"resFmtCode"`
	ReqFtmCode  string `json:"reqFtmCode"`
	ResDataProp string `json:"resDataProp"`
	ReqDataProp string `json:"reqDataProp"`
	ResExample  string `json:"resExample"`
	ReqExample  string `json:"reqExample"`
	FuncCode    string `json:"funcCode"`
	FuncName    string `json:"funcName"`
	GroupId     string `json:"groupId"`
	//add intercode as primary key
	InteCode     string `json:"inteCode"`
}
type UipInterFuncRequestJson struct {
	Com  string       `json:"com"`
	Data UipInterFunc `json:"data"`
}
type UipInterFuncsRequestJson struct {
	Com  string         `json:"com"`
	Data []UipInterFunc `json:"data"`
}
type SrcIntePseRequest struct {
	GroupId string `json:"groupId"`
	Fttext  string `json:"fttext"`
	Settext string `json:"settext"`
}

type BatchResQuest struct {
	GroupId string   `json:"groupId"`
	Data    []string `json:"data"`
}

//开放接口功能列表 成员
type List struct {
	OpenCode string `json:"openCode"`
	SrcName  string `json:"srcName"`
	SrcCode  string `json:"srcCode"`
}

//开放接口实体
type UipOpenInterface struct {
	Cmt        string `json:"cmt"`
	OperCode   string `json:"operCode"`
	Ctime      string `json:"cTime"`
	AfterCode  string `json:"afterCode"`
	SinteCode  string `json:"sinteCode"`
	BeforeCode string `json:"beforeCode"`
	FuncList   []List `json:"funcList"`
	ConnConf   string `json:"connConf"`
	InteType   string `json:"inteType"`
	InteCode   string `json:"inteCode"`
	InteName   string `json:"inteName"`
	GroupId    string `json:"groupId"`
}

type RequestOpenInterfaceJson struct {
	Com  string           `json:"com"`
	Data UipOpenInterface `json:"data"`
}
