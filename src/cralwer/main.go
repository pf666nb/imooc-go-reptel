package main

import (
	"cralwer/engine"
	"cralwer/zhenai/parser"
)

func main() {

	engine.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenhun,",
			ParserFunc: parser.ParseCityList,
		})
}
