package attached

import (
	"net/http"
	"encoding/json"
	"uip/common"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{
		Code: common.ErrorSystemErrId,
		Data: common.ErrorSystemErrMsg,
	}
}

func (r *Response) Answer(w http.ResponseWriter) error{
	jsonStr, err := json.Marshal(r)
	w.Write(jsonStr)
	return err
}
