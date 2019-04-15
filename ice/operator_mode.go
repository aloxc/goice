package ice

type OperatorMode byte

const (
	OperatorModeNormal      OperatorMode = 0
	OperatorModeNonmutating OperatorMode = 1
	OperatorModeIdempotent  OperatorMode = 2 //幂等信接口标示
)
