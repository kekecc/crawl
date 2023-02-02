package main

import (
	"reptile/handle"
	"reptile/info"
	"reptile/schedulers"
)

func main() {
	handler := handle.ConcurrentHandler{
		Scheduler: &schedulers.QueueScheduler{},
		Workers:   100,
	}
	handler.Run(handle.Request{
		Url:       "https://book.douban.com",
		ParseFunc: info.ParseContent,
	})
}
