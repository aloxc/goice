package ice

import "fmt"

//ice服务器侧发生异常，该异常是ice服务器端业务抛出的异常，不是ice框架的异常
type UserError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}

func NewUserError(address, operator, desc string, params interface{}) *UserError {
	return &UserError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}
func (this *UserError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

type TimeoutError struct {
	timeout  int
	address  string
	operator string
	params   interface{}
}

func NewTimeoutError(address, operator string, timeout int, params interface{}) *TimeoutError {
	return &TimeoutError{
		timeout:  timeout,
		address:  address,
		operator: operator,
		params:   params,
	}
}
func (this *TimeoutError) Error() string {
	return fmt.Sprintf("\nICE服务器端响应超时异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t超时时间: [%d]秒\n\n",
		this.address, this.operator, this.params, this.timeout)
}