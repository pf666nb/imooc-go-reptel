package main

import (
	"cralwer/engine"
	"cralwer/zhenai/parser"
)

func main() {

	engine.Run(
		engine.Request{
			Url:        "https://www.zhenai.com/zhenghun/",
			ParserFunc: parser.ParseCityList,
		})
}
