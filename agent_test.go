package gogetxueqiu

import (
	"fmt"
	"testing"
	"time"
)

func TestALL(t *testing.T) {
	// add a xueqiu.com account
	XueqiuAccounts["jxgzwd@163.com"] = "123456"//"E10ADC3949BA59ABBE56E057F20F883E" MD5
	
	// login
	_, err := Login()
	if err != nil {
		panic("login failed")
	}

	// stock real time
	stockrt, err := GetStockRT("SZ000625")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GetStockRT",*stockrt)
	
	// stock K list
	te := time.Now()
	tb := te.AddDate(0,-1,0)
	sKLPParams := &stockKListParams {
		symbol:"SZ000625",
		period:"1day",
		fuquanType:"before",
		begin:tb,
		end:te,
	}
	stockKPriceListHS, err := GetStockPriceListHS(*sKLPParams)
	fmt.Println("stockKListParams",*sKLPParams)
	fmt.Println("GetStockPriceListHS for 1 month:",stockKPriceListHS.Success,len(stockKPriceListHS.PriceListHS))
	
	// stock minutes list
	sMinsParams := &stockMinutesParams {
		symbol:"SZ000625",
		period:"1d",
		onemin:1,
	}
	stockPriceMins, err := GetStockPriceMinutes(*sMinsParams)
	fmt.Println("stockMinutesParams", *sMinsParams)
	fmt.Println("GetStockPriceMinutes for 1 day:",stockPriceMins.Success, len(stockPriceMins.PriceListMins))
}
