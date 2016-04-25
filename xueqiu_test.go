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
	sKLPParams := &StockKListParams {
		Symbol:"SZ000625",
		Period:"1day",
		FuquanType:"before",
		Begin:tb,
		End:te,
	}
	stockKPriceListHS, err := GetStockPriceListHS(*sKLPParams)
	fmt.Println("GetStockPriceListHS for 1 month:",stockKPriceListHS.Success,len(stockKPriceListHS.PriceListHS))
	
	// stock minutes list
	sMinsParams := &StockMinutesParams {
		Symbol:"SZ000625",
		Period:"1d",
		OneMin:1,
	}
	stockPriceMins, err := GetStockPriceMinutes(*sMinsParams)
	fmt.Println("GetStockPriceMinutes for 1 day:",stockPriceMins.Success, len(stockPriceMins.PriceListMins))
	
	// portfolio value list
	pValuesParams := &PfValuesParams {
		CubeSymbol:"ZH024581",
		Since:tb,
		Until:te,
	}
	pfValuesListHS, err := GetPfValueListHS(*pValuesParams)
	fmt.Println("GetPfValueListHS for 1 month:",pfValuesListHS.Name, len(pfValuesListHS.ListHS))
}
