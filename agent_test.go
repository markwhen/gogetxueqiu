package gogetxueqiu

import (
	"fmt"
	"testing"
)

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
	
	sLPParams := &stockListParams {
		symbol:"SZ000625",
		period:"1day",
		fuquanType:"before",
		begin:1423798115327,
		end:1429798115327,
	}
	StockPriceListHS, err := GetStockPriceListHS(*sLPParams)
	fmt.Println(*StockPriceListHS)
}
