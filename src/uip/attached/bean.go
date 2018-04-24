package attached

type UipDfmtMain struct {
	Cmt      string `json:"cmt"`
	OperCode string `json:"operCode"`
	Ctime    string `json:"ctime"`
	//dtlCode  string `json:"dtlCode"`
	ExtInfo  string `json:"extInfo"`
	FmtType  string `json:"fmtType"`
	FmtCode  string `json:"fmtCode"`
	FmtName  string `json:"fmtName"`
	GroupId  string `json:"groupId"`
}

type MainRequestJson struct {
	Com  string     `json:"com"`
	Data UipDfmtMain `json:"data"`
}

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

type SubRequestJson struct {
	Com  string     `json:"com"`
	Data UipDfmtSub `json:"data"`
}