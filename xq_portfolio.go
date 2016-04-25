package gogetxueqiu

import (
    "strconv"
    "errors"
    "encoding/json"
    "fmt"
)

// PfValueHS :
type PfValueHS struct {
    TimeStamp int64   `json:"time"`
    DateStr   string  `json:"date"`
    Value     float32 `json:"value"`
    Percent   float32 `json:"percent"`
}

// PfValueListHS :
type PfValueListHS struct {
    Symbol    string  `json:"symbol"`
    Name      string  `json:"name"`
    ListHS    []PfValueHS `json:"list"`
}

//GetPfValueListHS : get stock price(K) list in history
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