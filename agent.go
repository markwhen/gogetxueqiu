package gogetxueqiu

import (
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

var debugLogging = true

var myCookieJar *cookiejar.Jar

func getRequestUrl(urlBase string, params map[string]string) string {
	url := urlBase
	if params != nil && len(params) > 0 {
		url += "?"
		for k, v := range params {
			url += (k + "=" + v + "&")
		}
		url = url[:len(url)-1]
	}
	return url
}

// HTTPGet : request using GET method, will encode the get params into url
func HTTPGet(urlBase string, params map[string]string) (int, string, error) {
	url := getRequestUrl(urlBase, params)
	return httpRequest("GET", url, nil)
}

// HTTPGetBytes : request using GET method, will encode the get params into url
func HTTPGetBytes(urlBase string, params map[string]string) (int, []byte, error) {
	url := getRequestUrl(urlBase, params)
	return httpRequestBytes("GET", url, nil)
}

// HTTPGetJSON : request using GET method, will encode the get params into url
func HTTPGetJSON(urlBase string, params map[string]string) (int, *simplejson.Json, error) {
	url := getRequestUrl(urlBase, params)
	code, jsonBytes, err := httpRequestBytes("GET", url, nil)
	if code != 200 {
		return -1, nil, errors.New("return code: " + strconv.Itoa(code))
	}
	if err != nil {
		return -1, nil, err
	}
	json, err := simplejson.NewJson(jsonBytes)
	return 200, json, err
}

// HTTPPost : request using POST method, will encode the post params into formBody
func HTTPPost(urlBase string, params map[string]string) (int, string, error) {
	val := url.Values{}
	if params != nil && len(params) > 0 {
		for k, v := range params {
			val.Set(k, v)
		}
	}
	formBody := ioutil.NopCloser(strings.NewReader(val.Encode()))
	return httpRequest("POST", urlBase, formBody)
}

// httpRequest : default charset UTF-8
func httpRequest(method string, urlStr string, postBody io.ReadCloser) (int, string, error) {
	code, body, err := httpRequestBytes(method, urlStr, postBody)
	if err != nil {
		return -1, "", err
	}
	return code, string(body), nil
}

func httpRequestBytes(method string, urlStr string, postBody io.ReadCloser) (int, []byte, error) {
	if myCookieJar == nil {
		myCookieJar, _ = cookiejar.New(nil)
	}
	client := &http.Client{Jar: myCookieJar}
	if !((postBody == nil && method == "GET") || method == "POST") {
		panic("GET OR POST error")
	}
	req, err := http.NewRequest(method, urlStr, postBody)
	if err != nil {
		log.Println("NewRequest error")
		return -1, nil, err
	}
	if debugLogging {
		log.Println("httpRequest ", method, " ", urlStr, " ")
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Connection", "keep-alive")

	if postBody != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}
	// begin ..
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		if debugLogging {
			log.Println("do request error")
		}
		return -1, nil, err
	}
	// ungzip or not, and read to body
	var body []byte
	if res.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(res.Body)
		defer gzipReader.Close()
		if err != nil {
			return -1, nil, err
		}
		body, err = ioutil.ReadAll(gzipReader)
		if err != nil {
			return -1, nil, err
		}
	} else {
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return -1, nil, err
		}
	}
	return res.StatusCode, body, nil
}
