package main

import (
	"errors"
	"log"
	"strconv"
)

// XueqiuUrls : xueqiu urls map
var XueqiuUrls = map[string]string{
	"csrf":       "https://xueqiu.com/service/csrf",
	"login":      "https://xueqiu.com/user/login",
	"stock_curr": "https://xueqiu.com/v4/stock/quote.json",
}

// XueqiuAccounts : xueqiu accounts
var XueqiuAccounts = map[string]string{
	"jxgzwd@163.com": "E10ADC3949BA59ABBE56E057F20F883E", // 123456 MD5
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
			"password":    v,
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

//StockCurr : get stock current status
func StockCurr(stockStr string) (string, error) {
	code, res, err := HTTPGet(XueqiuUrls["stock_curr"], map[string]string{
		"code": stockStr,
	})
	if err != nil {
		log.Println("error when get ", XueqiuUrls["stock_curr"])
	}
	if code != 200 {
		return "", errors.New("code:" + strconv.Itoa(code))
	}
	return res, err
}
