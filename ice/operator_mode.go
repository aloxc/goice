package ice

type OperationMode byte

const (
	OperatorModeNormal      OperationMode = 0 //请求发送的标志
	OperatorModeNonmutating OperationMode = 1 //连接后发送的标注
	OperatorModeIdempotent  OperationMode = 2 //幂等信接口标示
)
