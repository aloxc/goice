package request

import "encoding/json"

type Request struct {
	Method string
	Params map[string]string
}
//构造用于execute方法的请求
func NewReqeust(method string,params map[string]string)*Request{
	return &Request{
		Method:method,
		Params:params,
	}
}
func (this*Request)String() string {
	r := make(map[string]interface{})
	r["method"] = this.Method
	r["params"] = this.Params
	if bytes, err := json.Marshal(r);err == nil{
		return string(bytes)
	}
	return ""
}