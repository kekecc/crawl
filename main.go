package main

import (
	"reptile/handle"
	"reptile/info"
)

func main() {
	handler := handle.ConcurrentHandler{
		Scheduler: handle.NormalScheduler{},
		Workers:   100,
	}
	handler.Run(handle.Request{
		Url:       "https://book.douban.com",
		ParseFunc: info.ParseContent,
	})
}
