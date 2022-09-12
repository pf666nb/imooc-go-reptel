package main

import (
	"cralwer/engine"
	"cralwer/scheduler"
	"cralwer/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(
		engine.Request{
			Url:        "https://www.zhenai.com/zhenghun/",
			ParserFunc: parser.ParseCityList,
		})
}
