package ice

import "fmt"

//ice服务器侧发生异常，该异常是ice服务器端业务抛出的异常，不是ice框架的异常
type UserUnknownError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}
type UserError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}
type ObjectNotExistsError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}
type FacetNotExistsError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}
type OperatorNotExistsError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}
type IceServerError struct {
	address  string
	operator string
	params   interface{}
	desc     string
}
type TimeoutError struct {
	timeout  int
	address  string
	operator string
	params   interface{}
}
type ConnectTimeoutError struct {
	timeout int
	address string
	network string
}

func NewTimeoutError(address, operator string, timeout int, params interface{}) *TimeoutError {
	return &TimeoutError{
		timeout:  timeout,
		address:  address,
		operator: operator,
		params:   params,
	}
}

func NewUserError(address, operator, desc string, params interface{}) *UserError {
	return &UserError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}
func NewUserUnknownError(address, operator, desc string, params interface{}) *UserUnknownError {
	return &UserUnknownError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}

func NewOperatorNotExistsError(address, operator, desc string, params interface{}) *OperatorNotExistsError {
	return &OperatorNotExistsError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}

func NewFacetNotExistsError(address, operator, desc string, params interface{}) *FacetNotExistsError {
	return &FacetNotExistsError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}

func NewIceServerError(address, operator, desc string, params interface{}) *IceServerError {
	return &IceServerError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}

func NewObjectNotExistsError(address, operator, desc string, params interface{}) *ObjectNotExistsError {
	return &ObjectNotExistsError{
		address:  address,
		operator: operator,
		params:   params,
		desc:     desc,
	}
}

func NewConnectTimeoutError(address, network string, timeout int) *ConnectTimeoutError {
	return &ConnectTimeoutError{
		address: address,
		network: network,
		timeout: timeout,
	}
}

func (this *UserUnknownError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

func (this *TimeoutError) Error() string {
	return fmt.Sprintf("\nICE服务器端响应超时异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t超时时间: [%d]秒\n\n",
		this.address, this.operator, this.params, this.timeout)
}

func (this *ObjectNotExistsError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

func (this *IceServerError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

func (this *FacetNotExistsError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

func (this *OperatorNotExistsError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

func (this *UserError) Error() string {
	return fmt.Sprintf("\nICE服务器端程序异常: \n\t地址: %s\n\t方法: %s\n\t参数: [%s]\n\t描述: \n\t%s\n\n", this.address, this.operator, this.params, this.desc)
}

func (this *ConnectTimeoutError) Error() string {
	return fmt.Sprintf("\n连接ICE服务器超时异常: [%s://%s?timeout=%d]\n", this.network, this.address, this.timeout)
}
