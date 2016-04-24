package gogetxueqiu

import (
	"fmt"
	"testing"
)

var debugLogging = false
var infoLogging = true

func TestALL(t *testing.T) {
	_, err := Login()
	if err != nil {
		panic("login failed")
	}

	stockrt, err := GetStockRT("SZ000625")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*stockrt)
}
