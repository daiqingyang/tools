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
var genPass bool
var resource_id string
var default_idx = 0 //默认使用配置文件中的第一个ak sk

func init() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	if default_idx >= len(config.CloudApis) {
		panic("cloud api index does not exist")
	}
	flag.BoolVar(&price, "p", false, "get price")
	flag.StringVar(&resource_id, "i", "", "resource_id")
	flag.BoolVar(&genPass, "r", false, "gen random password")
	flag.Parse()
}
func main() {
	if genPass {
		fmt.Println(tools.GenPassword())
	}
	GetPrice()
}
func GetPrice() {
	if price && resource_id != "" {
		cfg := config.CloudApis[default_idx]
		hw := tools.Hw{
			AK:        cfg.Ak,
			SK:        cfg.Sk,
			ProjectID: cfg.ProjectId,
		}
		hw.Init()
		// 查询月账单
		month, billTypeDesc, svcType, unitPrice, unit, err := hw.GetPriceByMonth(resource_id)
		if err != nil {
			panic(err)
		}
		// fmt.Println("资源id 				     月份    账单类型       云服务类型 		单价 单位")
		fmt.Println(resource_id, month, billTypeDesc, svcType, unitPrice, unit)
	}
}
