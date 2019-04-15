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
	ResponseType_Article
)
