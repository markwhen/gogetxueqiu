package gogetxueqiu

import (
	"fmt"
	"testing"
	"time"
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
	fmt.Println("GetStockRT",*stockrt)
	
	te := time.Now()
	tb := te.AddDate(0,-1,0)
	sLPParams := &stockListParams {
		symbol:"SZ000625",
		period:"1day",
		fuquanType:"before",
		begin:tb.Unix() * 1000,
		end:te.Unix() * 1000,
	}
	StockPriceListHS, err := GetStockPriceListHS(*sLPParams)
	fmt.Println("stockListParams",*sLPParams)
	fmt.Println("GetStockPriceListHS for 1 month",*StockPriceListHS)
}
