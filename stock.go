package gogetxueqiu

import (
    "strconv"
    "reflect"
    "errors"
    "encoding/json"
    "fmt"
    //"github.com/bitly/go-simplejson"
)

var stringNameMap = map[string]string {
    "Symbol":"symbol", "Exchange":"exchange", "Code":"code", "Name":"name", "CurrencyUnit":"currency_unit",
}

var uint64NameMap = map[string]string {
     "TotalShares":"totalShares","UpdateBasicAt":"updateAt", "UpdateAt":"updateAt",
}

var float32NameMap = map[string]string {
    "Current":"current", "Percentage":"percentage", "Change":"change", "Open":"open", "Close":"close",
    "LastClose": "last_close", "High":"high", "Low":"low", "MarketCapital":"marketCapital",
    "RiseStop":"rise_stop", "FallStop":"fall_stop", "PELYR":"pe_lyr", "PETTM":"pe_ttm",
    "EPS":"eps", "PSR":"psr", "PB":"pb", "Divident":"dividend", "Volume":"volume",
}

// StockBasic : basic info from real time info
type StockBasic struct {
    Symbol          string  `json:"symbol"`
    Exchange        string  `json:"exchange"`
    Code            string  `json:"code"`
    Name            string  `json:"name"`
    CurrencyUnit    string  `json:"currency_unit"`
    TotalShares     uint64  `json:"totalShares"`
    UpdateBasicAt   uint64  `json:"updateAt"`
}

// StockPriceRT : real time price for now
type StockPriceRT struct {
    Current         float32  `json:"current"`
    Percentage      float32  `json:"percentage"`
    Change          float32  `json:"change"`
    Open            float32  `json:"open"`
    Close           float32  `json:"close"`
    LastClose       float32  `json:"last_close"`
    High            float32  `json:"high"`
    Low             float32  `json:"low"`
    MarketCapital   float32  `json:"marketCapital"`
    RiseStop        float32  `json:"rise_stop"`
    FallStop        float32  `json:"fall_stop"`
    Volume          float32  `json:"volume"`
    PELYR           float32  `json:"pe_lyr"`
    PETTM           float32  `json:"pe_ttm"`
    EPS             float32  `json:"eps"`
    PSR             float32  `json:"psr"`
    PB              float32  `json:"pb"`
    Divident        float32  `json:"dividend"`
    UpdateAt        uint64   `json:"updateAt"`
}

// StockRT : Stock RealTime info
type StockRT struct {
    StockBasic
    StockPriceRT
}

// StockPriceHS : Stock Price(K) in HiStory
type StockPriceHS struct {
    Volume      uint64   `json:"volume"`
    Turnrate    float32  `json:"turnrate"`
    Open        float32  `json:"open"`
    Close       float32  `json:"close"`
    High        float32  `json:"high"`
    Low         float32  `json:"low"`
    Change      float32  `json:"chg"`
    Percentage  float32  `json:"percent"`
    MA5         float32  `json:"ma5"`
    MA10        float32  `json:"ma10"`
    MA20        float32  `json:"ma20"`
    MA30        float32  `json:"ma30"`
    MACD        float32  `json:"macd"`
    DEA         float32  `json:"dea"`
    DIF         float32  `json:"dif"`
    Time        string   `json:"time"`
}

// StockPriceListHS : contains K price list
type StockPriceListHS struct {
    Success     string    `json:"success"`
    PriceListHS []StockPriceHS  `json:"chartlist"`
}

// StockPriceMin : Stock Price(minutes in recent day) in HiStory
type StockPriceMin struct {
    Volume   uint64 `json:"volume"`
    AvgPrice float32 `json:"avg_price"`
    Current  float32 `json:"current"`
    Time     string  `json:"time"`
}
// StockPriceMins : contains the minute price list
type StockPriceMins struct {
    Success       string    `json:"success"`
    PriceListMins []StockPriceMin  `json:"chartlist"`
}

//GetStockPriceListHS : get stock price(K) list in history
func GetStockPriceListHS(reqParams stockKListParams) (*StockPriceListHS, error) {
	code, res, err := HTTPGetBytes(XueqiuUrls["stock_k_list"], map[string]string{
        "symbol": reqParams.symbol,
        "period": reqParams.period,
        "type": reqParams.fuquanType,
        "begin": strconv.FormatInt(reqParams.begin.Unix() * 1000, 10),
        "end": strconv.FormatInt(reqParams.end.Unix() * 1000, 10),
    })
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("code:" + strconv.Itoa(code))
	}
    stockPLHS := new(StockPriceListHS)
    err = json.Unmarshal(res, stockPLHS)
    if err != nil {
        fmt.Println("GetStockPriceListHS err:", err)
        return nil, err
    }
    return stockPLHS, nil
}


//GetStockPriceMinutes : get stock price minutes list in a day
func GetStockPriceMinutes(reqParams stockMinutesParams) (*StockPriceMins, error) {
	code, res, err := HTTPGetBytes(XueqiuUrls["stock_minutes"], map[string]string{
        "symbol": reqParams.symbol,
        "period": reqParams.period,
        "one_min": strconv.FormatInt(int64(reqParams.onemin),10),
    })
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("code:" + strconv.Itoa(code))
	}
    stockPMins := new(StockPriceMins)
    err = json.Unmarshal(res, stockPMins)
    if err != nil {
        fmt.Println("GetStockPriceMinutes err:", err)
        return nil, err
    }
    return stockPMins, nil
}

// fromMap : 
// get StockRT data from json.Map(), becuase xueqiu don't give a standard format, so parse it by reflection
func (stockrt *StockRT) fromMap(mp map[string]interface{}) error {
    for k,v := range stringNameMap {
        if mp[v] == nil {
            continue
        }
        if valuStr, ok := mp[v].(string); ok {
            reflect.ValueOf(stockrt).Elem().FieldByName(k).SetString(valuStr)
        }
    }
    for k,v := range uint64NameMap {
        if mp[v] == nil {
            continue
        }
        if valuStr, ok := mp[v].(string); ok {
            valu64, err := strconv.ParseUint(valuStr, 10, 64)
            if err != nil {
                return err
            }
            reflect.ValueOf(stockrt).Elem().FieldByName(k).SetUint(valu64)
        }
        
    }
    for k,v := range float32NameMap {
        if mp[v] == nil {
            continue
        }
        if valuStr, ok := mp[v].(string); ok {
            valu64, err := strconv.ParseFloat(valuStr, 64)
            if err != nil {
                return err
            }
            reflect.ValueOf(stockrt).Elem().FieldByName(k).SetFloat(valu64)
        }
    }
    return nil
}

//GetStockRT : get stock current status
func GetStockRT(stockStr string) (*StockRT, error) {
	code, res, err := HTTPGetJSON(XueqiuUrls["stock_rt"], map[string]string{"code": stockStr})
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("code:" + strconv.Itoa(code))
	}
	stockrt := new(StockRT)
	err = stockrt.fromMap(res.Get(stockStr).MustMap())
	return stockrt, err
}