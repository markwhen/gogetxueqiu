package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/ynsfsmj/gogetxueqiu"
)

var wg sync.WaitGroup

func fetchAndPrint(stockCode string, dataTube chan string) (string, error) {
	stockrt, err := gogetxueqiu.GetStockRT(stockCode)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	out, err := json.Marshal(*stockrt)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(out), nil
}

func worker(codeTube <-chan string, dataTube chan string) {
	var code string

	for {
		code = <-codeTube
		fmt.Println("Begin to get ", code)
		res, err := fetchAndPrint(code, dataTube)
		if err != nil {
			wg.Done()
			continue
		}
		dataTube <- string(res)
	}
}

func speaker(dataTube chan string) {
	var cnt int
	for {
		out := <-dataTube
		cnt++
		fmt.Printf("Get Data[%d] : %s\n", cnt, out)
		wg.Done()
	}
}

func main() {
	gogetxueqiu.XueqiuAccounts["jxgzwd@163.com"] = "123456" //"E10ADC3949BA59ABBE56E057F20F883E" MD5

	// login
	_, err := gogetxueqiu.Login()
	if err != nil {
		panic("login failed")
	}

	// open data file
	file, err := os.Open("codes.log")
	if err != nil {
		panic("read failed")
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	var codes = make(chan string, 20)
	var datas = make(chan string)

	go speaker(datas) // init data speaker

	// start workers
	for i := 0; i < 20; i++ {
		go worker(codes, datas)
	}

	// send tasks
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("read failed")
		}
		wg.Add(1)
		codes <- "SH" + strings.TrimSpace(line)
	}

	// wait to finish
	wg.Wait()

}
