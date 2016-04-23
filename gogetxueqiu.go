package main

import (
	"fmt"
)

var debugLogging = false
var infoLogging = true

func main() {
	username, err := Login()
	if err != nil {
		panic("login failed")
	}
	fmt.Println(username)
	content, err := StockCurr("SZ000625")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
