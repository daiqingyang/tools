package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/daiqingyang/tools"
	"gopkg.in/yaml.v3"
)

var config Config
var price bool
var month string
var resource_id string

func init() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	flag.BoolVar(&price, "p", false, "get price")
	flag.StringVar(&month, "m", "2025-07", "查询的资源消费记录所在账期,格式:YYYY-MM")
	flag.StringVar(&resource_id, "i", "", "resource_id")
	flag.Parse()
}
func main() {
	GetPrice()
}
func GetPrice() {
	if price && resource_id != "" {
		hw := tools.Hw{
			AK: config.CloudApis[0].Ak,
			SK: config.CloudApis[0].Sk,
		}
		hw.Init()
		// 查询月账单
		month, billTypeDesc, svcType, unitPrice, unit, err := hw.GetPriceByMonth(resource_id)
		if err != nil {
			panic(err)
		}
		fmt.Println(month, billTypeDesc, svcType, unitPrice, unit)
	}
}
