package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/daiqingyang/tools"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var config Config
var getPrice bool
var genPass bool
var checkPass bool
var resource_id string
var runWeb bool
var runSomeTest bool
var default_idx = 0 //默认使用配置文件中的第一个ak sk
var log *logrus.Logger

func init() {
	log = logrus.New()
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
	flag.BoolVar(&getPrice, "p", false, "get price")
	flag.StringVar(&resource_id, "i", "", "resource_id")
	flag.BoolVar(&genPass, "r", false, "gen random password")
	flag.BoolVar(&checkPass, "c", false, "check password")
	flag.BoolVar(&runWeb, "l", false, "run Web")
	flag.BoolVar(&runSomeTest, "t", false, "run some test ")
	flag.Parse()
}
func main() {
	if genPass {
		fmt.Println(tools.GenPassword())
	} else if getPrice {
		GetPrice()
	} else if checkPass {
		examples := []string{
			"aaaaaa",
			"123456",
			"password",
			"Bocom_abdx",
			"AKIAIOSFODNN7EXAMPLE",
			"AKIAIOSFODNN7Ea1MPLE",
			"Aq8XnP9tR4v1ZkLm",
			"qLc7D9k5ut",
			"92c7a9b1065af0b2d8fdd7919acbdcde", //md5
			"U2FsdGVkX1+M2wNVUgsTsGd8kd4UKF7OOxLjn4THcQc=",
			"U2FsdGVkX1/bhXgKR8WNnCijsRkDtOg5IA4XKLo02rM=",
			"U2FsdGVkX1/GW0a0LbSJWDx8dhnrP2Towd0=",
			"U2FsdGVkX1+lez4GHCeUN/+EQjH4V/pccWE=",
			"U2FsdGVkX18b049FphglQyapKg44s6caNCng8k4O/14=",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", // JWT token 片段
		}

		for _, str := range examples {
			fmt.Printf("字符串: %-40s 熵值: %10.3f\n", str, tools.ShannonEntropy(str))
		}
	} else if runWeb {
		RunWeb()
	} else if runSomeTest {
		RunOrmTest()
	}
}
func GetPrice() {
	if resource_id != "" {
		cfg := config.CloudApis[default_idx]
		//华为公有云
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
