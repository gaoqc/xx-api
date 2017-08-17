package controllers

type RetVO struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessVO(d interface{}) RetVO {
	return RetVO{Code: SuccCode, Msg: SuccMsg, Data: d}
}
func GetRetVO(code, msg string, d interface{}) RetVO {
	return RetVO{Code: code, Msg: msg, Data: d}
}
