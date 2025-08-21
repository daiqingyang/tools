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
	flag.StringVar(&resource_id, "i", "", "resource_id")
	flag.Parse()
}
func main() {
	if price && resource_id != "" {
		hw := tools.Hw{
			AK: config.CloudApis[0].Ak,
			SK: config.CloudApis[0].Sk,
		}
		hw.Init()
		official_price, real_price, err := hw.GetPrice(resource_id)
		if err != nil {
			panic(err)
		}
		fmt.Println(official_price, real_price)
	}
}
