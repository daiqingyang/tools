package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/daiqingyang/tools/core"
)

type Hw struct {
	AK string
	SK string
	s  core.Signer
}

func (hw *Hw) Init() {
	hw.s = core.Signer{
		Key:    hw.AK,
		Secret: hw.SK,
	}
}
func (hw *Hw) GetPrice(resource_id string) (office_price, real_price float64, err error) {
	url := "https://bss.myhuaweicloud.com/v2/bills/customer-bills/res-fee-records?cycle=2025-07&resource_id=" + resource_id
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}
	hw.s.Sign(req) //进行签名，执行此函数会在请求中添加用于签名的X-Sdk-Date头和Authorization头
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	var rst RespFeeRecords
	err = json.Unmarshal(body, &rst)
	if err != nil {
		return
	}
	if len(rst.FeeRecords) > 0 {
		office_price = rst.FeeRecords[0].OfficialAmount
		real_price = rst.FeeRecords[0].Amount
	}
	return
}
