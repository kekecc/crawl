package main

import (
	datasave "reptile/dataSave"
	"reptile/handle"
	"reptile/info"
	"reptile/schedulers"
)

func main() {
	handler := handle.ConcurrentHandler{
		Scheduler: &schedulers.QueueScheduler{},
		Workers:   100,
		DataSave:  datasave.Save(),
	}
	handler.Run(handle.Request{
		Url:       "https://book.douban.com",
		ParseFunc: info.ParseContent,
	})
}
