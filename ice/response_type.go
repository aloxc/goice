package ice

type ResponseType byte

const (
	ResponseType_Void ResponseType = iota
	ResponseType_String
	ResponseType_Int
	ResponseType_Double
	ResponseType_Float
)
