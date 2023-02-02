package handle

import (
	"log"
	"reptile/fetch"
)

type Scheduler interface {
	Submit(Request)
	//GetChan(chan Request)
	GetWorker(chan Request)
	Run()
}

type ConcurrentHandler struct {
	Scheduler Scheduler
	Workers   int
}

func (ch *ConcurrentHandler) Run(urls ...Request) {
	//tasks := make(chan Request)
	results := make(chan ParseRes)
	ch.Scheduler.Run()
	//ch.Scheduler.GetChan(tasks)
	for i := 0; i < ch.Workers; i++ {
		CreateWorker(results, ch.Scheduler)
	}

	for _, url := range urls {
		ch.Scheduler.Submit(url)
	}

	for {
		res := <-results

		for _, content := range res.Contents {
			log.Println("get content: ", string(content.([]byte)))
		}

		for _, resquest := range res.Requests {
			ch.Scheduler.Submit(resquest)
		}
	}
}

func CreateWorker(results chan ParseRes, s Scheduler) { //工人
	tasks := make(chan Request)
	go func() {
		for {
			s.GetWorker(tasks)
			req := <-tasks
			res, err := HandleTask(req)
			if err != nil {
				log.Println("handle error: ", err)
				continue
			}
			results <- res
		}
	}()
}

func HandleTask(req Request) (ParseRes, error) {
	log.Println("fetch url:", req.Url)
	body, err := fetch.Fetch(req.Url)
	if err != nil {
		log.Println("fetch err: ", err)
		return ParseRes{}, err
	}
	result := req.ParseFunc(body) //会有其他任务
	return result, nil
}
