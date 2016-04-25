package gogetxueqiu

import (
    "strconv"
    "errors"
    "encoding/json"
    "fmt"
    "time"
)

// PfValueHS : a net value item in a list of a portfolio
type PfValueHS struct {
    TimeStamp int64   `json:"time"`
    //DateStr   string  `json:"date"`
    Value     float32 `json:"value"`
    Percent   float32 `json:"percent"`
}

// PfValueListHS : the net values list of a portfolio
type PfValueListHS struct {
    Symbol    string  `json:"symbol"`
    Name      string  `json:"name"`
    ListHS    []PfValueHS `json:"list"`
}

// PfBasic : the basic info of a portfolio from serveral data source, get from pf_scores
type PfBasic struct {
    Symbol    string  `json:"symbol"`
    Name      string  `json:"name"`     // get from nav_daily
    Market    string  `json:"market"`
    ID        int64   `json:"cube_id"`
}

// PfRebalanceStock : a stock rebalancing of a record 
type PfRebalanceStock struct {
    StockName   string  `json:"stock_name"`
    StockSymbol string  `json:"stock_symbol"`
    Weight      float32  `json:"weight"`
    TargetWeight float32 `json:"target_weight"`
    Price       float32  `json:"price"`
    UpdatedAt   int64   `json:"updated_at"`
}

// PfRebalanceHS : an item of rebalance record of a portfolio
type PfRebalanceHS struct {
    Status    string  `json:"status"` //success, failed, canceled
    CashValue float64 `json:"cash_value"`
    RebalancingHistories []PfRebalanceStock `json:"rebalancing_histories"`
    UpdatedAt int64   `json:"updated_at"`
}

// PfReBalanceListPage : a page of rebalance records of a portfolio
type PfReBalanceListPage struct {
    Count     int64   `json:"count"`
    Page      int64   `json:"page"`
    TotalCount int64  `json:"totalCount"`
    PageList  []PfRebalanceHS  `json:"list"`
    MaxPage   int64  `json:"maxPage"`
}

// GetPfBasic : get portfolio basic infomation
func GetPfBasic(symbol string) (*PfBasic, error) {
    code, res, err := HTTPGetBytes(XueqiuUrls["pf_scores"], map[string]string {
        "symbol": symbol,
        "ua"    : "app",
    })
    if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("code:" + strconv.Itoa(code))
	}
    pfBasic := new(PfBasic)
    err = json.Unmarshal(res, &pfBasic)
    if err != nil {
        fmt.Println("GetPfBasic", err)
        return nil, err
    }
    // get name from nav_daily api
    pValuesParams := &PfValuesParams {
		CubeSymbol:symbol,
		Since:time.Now(),
		Until:time.Now(),
	}
	pfValuesListHS, err := GetPfValueListHS(*pValuesParams)
	if err != nil {
		fmt.Println("GetPfValueListHS get Name error", err) // get Name error
	} else {
        pfBasic.Name = pfValuesListHS.Name
	}
    
    return pfBasic, nil
}

// GetPfValueListHS : get portfolio value list in history
func GetPfValueListHS(reqParams PfValuesParams) (*PfValueListHS, error) {
	code, res, err := HTTPGetBytes(XueqiuUrls["pf_daily"], map[string]string{
        "cube_symbol": reqParams.CubeSymbol,
        "since": strconv.FormatInt(reqParams.Since.Unix() * 1000, 10),
        "until": strconv.FormatInt(reqParams.Until.Unix() * 1000, 10),
    })
    //fmt.Println(string(res))
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("code:" + strconv.Itoa(code))
	}
    var pfVLHS []PfValueListHS
    err = json.Unmarshal(res, &pfVLHS)
    if err != nil {
        fmt.Println("GetPfValueListHS err:", err)
        return nil, err
    }
    return &pfVLHS[0], nil
}

// GetPfRebalanceListPage : get portfolio rebalance action list in history, paged
func GetPfRebalanceListPage(reqParams PfRebalanceParams) (*PfReBalanceListPage, error) {
    // if params exceeds the range, it will return 400
	code, res, err := HTTPGetBytes(XueqiuUrls["pf_rebalance"], map[string]string{
        "cube_symbol": reqParams.CubeSymbol,
        "page" : strconv.FormatInt(reqParams.Page, 10),
        "count": strconv.FormatInt(reqParams.Count, 10),
    })
    //fmt.Println(string(res))
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("code:" + strconv.Itoa(code))
	}
    pfRBLP := new(PfReBalanceListPage)
    err = json.Unmarshal(res, pfRBLP)
    if err != nil {
        fmt.Println("GetPfRebalanceListPage err:", err)
        return nil, err
    }
    return pfRBLP, nil
}