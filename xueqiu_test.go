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

	// stock real time info
	stockrt, err := GetStockRT("SZ000625")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetStockRT",*stockrt)
	}
	
	te := time.Now()
	tb := te.AddDate(0,-1,0) // a month ago
	
	// stock K list
	sKLPParams := &StockKListParams {
		Symbol:"SZ000625",
		Period:"1day",
		FuquanType:"before",
		Begin:tb,
		End:te,
	}
	stockKPriceListHS, err := GetStockPriceListHS(*sKLPParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetStockPriceListHS for 1 month:",stockKPriceListHS.Success,len(stockKPriceListHS.PriceListHS))
	}
	
	// stock minutes list in the recent day
	sMinsParams := &StockMinutesParams {
		Symbol:"SZ000625",
		Period:"1d",
		OneMin:1,
	}
	stockPriceMins, err := GetStockPriceMinutes(*sMinsParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetStockPriceMinutes for 1 day:",stockPriceMins.Success, len(stockPriceMins.PriceListMins)) // max 242
	}
	
	// portfolio value list
	pValuesParams := &PfValuesParams {
		CubeSymbol:"ZH024581",
		Since:tb,
		Until:te,
	}
	pfValuesListHS, err := GetPfValueListHS(*pValuesParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPfValueListHS for 1 month:",pfValuesListHS.Name, len(pfValuesListHS.ListHS))
	}
	
	// portfolio rebalancing list
	pRebalanceParams := &PfRebalanceParams {
		CubeSymbol:"ZH024581",
		Count:50,
		Page: 1,
	}
	pRebalanceListPage, err := GetPfRebalanceListPage(*pRebalanceParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPfRebalanceListPage:", pRebalanceListPage.Count, len(pRebalanceListPage.PageList))
	}
	
	// portfolio basic information
	pBasic, err := GetPfBasic("ZH024581")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPfBasic:", pBasic.Name, pBasic.ID)
	}
}
