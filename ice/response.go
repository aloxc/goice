package ice

type ResponseType byte

const (
	ResponseType_Void ResponseType = iota
	ResponseType_String
	ResponseType_Bool
	ResponseType_Int8
	ResponseType_Int16
	ResponseType_Int
	ResponseType_Int64
	ResponseType_Float32
	ResponseType_Float64
	ResponseType_Execute
)
//响应结果(此结果是ice接口方法执行后返回的结果，而不是网络请求后返回的结果)
type Response struct {
	Code int8
	Message string
	Data interface{}
}