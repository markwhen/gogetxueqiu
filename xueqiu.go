package gogetxueqiu

import (
	"errors"
	"log"
	"time"
	"crypto/md5"
	"encoding/hex"
	//"strconv"
)

// XueqiuAccounts : xueqiu accounts
var XueqiuAccounts = map[string]string{}

// XueqiuUrls : xueqiu urls map
var XueqiuUrls = map[string]string{
	// login
	"csrf":			"https://xueqiu.com/service/csrf",
					//api
	"login":		"https://xueqiu.com/user/login",
					//username, password, areacode, remember_me
	// stock
	"stock_rt":		"https://xueqiu.com/v4/stock/quote.json",
					//code
	"stock_k_list":	"https://xueqiu.com/stock/forchartk/stocklist.json",
					//symbol, period(1day, 1week, 1month), type(before, after), begin, end (timestamp micsec)
	"stock_minutes":"https://xueqiu.com/stock/forchart/stocklist.json",
					//symbol, period(1d), one_min(1, 2)
	"stock_buysell":"https://xueqiu.com/stock/pankou.json",
					//symbol
	"stocks_price" :"https://xueqiu.com/stock/quotep.json",
					//stockid

	// portfolio
	"pf_daily":		  "https://xueqiu.com/cubes/nav_daily/all.json",
					  //cube_symbol, since, until (timestamp micsec)
	"pf_recommend":	  "https://xueqiu.com/cubes/discover/rank/cube/list.json?category=14",
					  //fixed catefory
	"pf_rank_percent":"https://xueqiu.com/cubes/data/rank_percent.json",
					  //cube_id, market(cn), dimension(annual)
	"pf_rebalance":   "https://xueqiu.com/cubes/rebalancing/history.json",
					  //cube_symbol, page, count, since, until
	"pf_scores":      "https://xueqiu.com/cubes/rank/summary.json",
					  //symbol, ua(app)
}

// StockKListParams :
type StockKListParams struct {
	Symbol string
	Period string
	FuquanType string
	Begin time.Time
	End time.Time
}

// StockMinutesParams :
type StockMinutesParams struct {
	Symbol string
	Period string
	OneMin int
}

// PfValuesParams :
type PfValuesParams struct {
	CubeSymbol string
	Since time.Time
	Until time.Time
}

// PfRebalanceParams :
type PfRebalanceParams struct {
	CubeSymbol string
	Count int64	// count <= 50
	Page  int64	// from 1 to 20
}

// GetMd5HexStr : calculate MD5
func GetMd5HexStr(str string) string {
	md5Ctx := md5.New()
    md5Ctx.Write([]byte(str))
    cipherStr := md5Ctx.Sum(nil)
    return hex.EncodeToString(cipherStr)
}

// Login : return ok, username
func Login() (string, error) {
	code, _, err := HTTPGet(XueqiuUrls["csrf"], map[string]string{
		"api": "%2Fuser%2Flogin",
	})
	if err != nil {
		return "", err
	}
	if code != 200 {
		return "", errors.New("Loin CSRF failed")
	}
	for k, v := range XueqiuAccounts {
		code, _, err := HTTPPost(XueqiuUrls["login"], map[string]string{
			"username":    k,
			"password":    GetMd5HexStr(v),
			"remember_me": "on",
			"areacode":    "86",
		})
		if err != nil {
			return "", err
		} else if err == nil && code == 200 {
			log.Println("Login with username", k)
			return k, nil
		}
	}
	return "", errors.New("Login Failed")
}
