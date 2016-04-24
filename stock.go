package gogetxueqiu

import (
    "strconv"
    "reflect"
    "errors"
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
    "RiseStop":"rise_stop", "FallStop":"fall_stop", "Volume":"volume", "PELYR":"pe_lyr", "PETTM":"pe_ttm",
    "EPS":"eps", "PSR":"psr", "PB":"pb", "Divident":"dividend",
}

// StockBasic :
type StockBasic struct {
    Symbol string
    Exchange string
    Code string
    Name string
    CurrencyUnit string
    TotalShares uint64
    UpdateBasicAt uint64
}

// StockPriceRT :
type StockPriceRT struct {
    Current float32
    Percentage float32
    Change float32
    Open float32
    Close float32
    LastClose float32
    High float32
    Low float32
    MarketCapital float32
    RiseStop float32
    FallStop float32
    Volume float32
    PELYR float32
    PETTM float32
    EPS float32
    PSR float32
    PB float32
    Divident float32
    UpdateAt uint64
}

// StockRT : Stock RealTime info
type StockRT struct {
    StockBasic
    StockPriceRT
}

// fromMap : get StockRT data from json.Map()
func (stockrt *StockRT) fromMap(mp map[string]interface{}) error {
    for k,v := range stringNameMap {
        if mp[v] == nil {
            return errors.New("KEY " + v + " not exists")
        }
        if valuStr, ok := mp[v].(string); ok {
            reflect.ValueOf(stockrt).Elem().FieldByName(k).SetString(valuStr)
        }
    }
    for k,v := range uint64NameMap {
        if mp[v] == nil {
            return errors.New("KEY " + v + " not exists")
        }
        valu64, err := strconv.ParseUint(mp[v].(string), 10, 64)
        if err != nil {
            valu64 = 0
        }
        reflect.ValueOf(stockrt).Elem().FieldByName(k).SetUint(valu64)
    }
    for k,v := range float32NameMap {
        if mp[v] == nil {
            return errors.New("KEY " + v + " not exists")
        }
        valu64, err := strconv.ParseFloat(mp[v].(string), 64)
        if err != nil {
            valu64 = 0
        }
        reflect.ValueOf(stockrt).Elem().FieldByName(k).SetFloat(valu64)
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