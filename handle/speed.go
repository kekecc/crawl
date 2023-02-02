package handle

import (
	"log"
	"reptile/fetch"
)

type Scheduler interface {
	Submit(Request)
	GetChan(chan Request)
}

type ConcurrentHandler struct {
	Scheduler NormalScheduler
	Workers   int
}

type NormalScheduler struct {
	requests chan Request
}

func (ns *NormalScheduler) Submit(req Request) {
	//ns.requests <- req 出现死锁
	go func() {
		ns.requests <- req
	}()
}

func (ns *NormalScheduler) GetChan(requests chan Request) {
	ns.requests = requests
}

func (ch *ConcurrentHandler) Run(urls ...Request) {
	tasks := make(chan Request)
	results := make(chan ParseRes)

	ch.Scheduler.GetChan(tasks)
	for i := 0; i < ch.Workers; i++ {
		CreateWorker(tasks, results)
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

func CreateWorker(tasks chan Request, results chan ParseRes) { //工人
	go func() {
		for {
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
