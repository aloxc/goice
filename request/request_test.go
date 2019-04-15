package request

import (
	"fmt"
	"testing"
)

func TestRequest_String(t *testing.T) {
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "15"
	req := NewReqeust(method, params)
	fmt.Println(req.String())
}
func TestNewReqeust(t *testing.T) {
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "15"
	NewReqeust(method,params)
}
