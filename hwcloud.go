package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/daiqingyang/tools/core"
)

var billTypeMap = map[int]string{
	1:   "消费-新购",
	2:   "消费-续订",
	3:   "消费-变更",
	5:   "消费-使用",
	8:   "消费-自动续订",
	12:  "消费-按时计费",
	18:  "消费-按月付费",
	103: "消费-按年付费",
}

// bill_type
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

// 根据月度账单，获取价格
// month=YYYY-MM
func (hw *Hw) GetPriceByMonth(resource_id string) (month, billTypeDesc, svcType string, unitPrice float64, unit string, err error) {
	now := time.Now()
	var count = 0
	for count < 24 {
		month = now.Format("2006-01")

		url := "https://bss.myhuaweicloud.com/v2/bills/customer-bills/res-fee-records?cycle=" + month + "&resource_id=" + resource_id
		method := "GET"

		payload := strings.NewReader(``)

		client := &http.Client{}
		var req *http.Request
		req, err = http.NewRequest(method, url, payload)
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
			rr := rst.FeeRecords[0]
			// spew.Dump(rst.FeeRecords[0])
			// 账单类型：bill_type
			// 1：消费-新购
			// 2：消费-续订
			// 3：消费-变更
			// 5：消费-使用
			// 8：消费-自动续订
			// 12：消费-按时计费
			// 18：消费-按月付费
			// 103：消费-按年付费
			bt := []int{1, 2, 3, 5, 8, 12, 18, 103}
			if ContainsInt(bt, rr.BillType) {
				billTypeDesc = billTypeMap[rr.BillType]
				svcType = rr.CloudServiceType
				unitPrice = rr.UnitPrice
				unit = rr.Unit

				break
			}
		}
		now = now.AddDate(0, -1, 0) //往前一个月
		count++
	}
	return
}

// to do
// 未实现，不能使用
func (hw *Hw) GetPriceByProduct(args ProductReqInfo) (err error) {

	url := "https://bss.myhuaweicloud.com/v2/bills/ratings/period-resources/subscribe-rate"
	method := "POST"
	var b []byte
	b, err = json.Marshal(args)
	if err != nil {
		return
	}
	payload := bytes.NewBuffer(b)

	client := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		return
	}
	hw.s.Sign(req) //进行签名，执行此函数会在请求中添加用于签名的X-Sdk-Date头和Authorization头

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	return
}
