package main

import (
	"fmt"
	"time"

	"github.com/ynsfsmj/gogetxueqiu"
)

func main() {
	// add a xueqiu.com account
	gogetxueqiu.XueqiuAccounts["jxgzwd@163.com"] = "123456" //"E10ADC3949BA59ABBE56E057F20F883E" MD5

	// login
	_, err := gogetxueqiu.Login()
	if err != nil {
		panic("login failed")
	}

	// stock real time info
	stockrt, err := gogetxueqiu.GetStockRT("SH601166")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetStockRT", *stockrt)
	}

	te := time.Now()
	tb := te.AddDate(0, -1, 0) // a month ago

	// stock K list
	sKLPParams := &gogetxueqiu.StockKListParams{
		Symbol:     "SZ000625",
		Period:     "1day",
		FuquanType: "before",
		Begin:      tb,
		End:        te,
	}
	stockKPriceListHS, err := gogetxueqiu.GetStockPriceListHS(*sKLPParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetStockPriceListHS for 1 month:", stockKPriceListHS.Success, len(stockKPriceListHS.PriceListHS))
	}

	// stock minutes list in the recent day
	sMinsParams := &gogetxueqiu.StockMinutesParams{
		Symbol: "SZ000625",
		Period: "1d",
		OneMin: 1,
	}
	stockPriceMins, err := gogetxueqiu.GetStockPriceMinutes(*sMinsParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetStockPriceMinutes for 1 day:", stockPriceMins.Success, len(stockPriceMins.PriceListMins)) // max 242
	}

	// portfolio value list
	pValuesParams := &gogetxueqiu.PfValuesParams{
		CubeSymbol: "ZH024581",
		Since:      tb,
		Until:      te,
	}
	pfValuesListHS, err := gogetxueqiu.GetPfValueListHS(*pValuesParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPfValueListHS for 1 month:", pfValuesListHS.Name, len(pfValuesListHS.ListHS))
	}

	// portfolio rebalancing list
	pRebalanceParams := &gogetxueqiu.PfRebalanceParams{
		CubeSymbol: "ZH024581",
		Count:      50,
		Page:       1,
	}
	pRebalanceListPage, err := gogetxueqiu.GetPfRebalanceListPage(*pRebalanceParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPfRebalanceListPage:", pRebalanceListPage.Count, len(pRebalanceListPage.PageList))
	}

	// portfolio basic information
	pBasic, err := gogetxueqiu.GetPfBasic("ZH024581")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPfBasic:", pBasic.Name, pBasic.ID)
	}
}
