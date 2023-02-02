package schedulers

import "reptile/handle"

type QueueScheduler struct {
	requests chan handle.Request
	workers  chan chan handle.Request //工人 每个工人对应一个任务池
}

func (qs *QueueScheduler) Worker() chan handle.Request {
	return make(chan handle.Request)
}

func (qs *QueueScheduler) Submit(req handle.Request) {
	qs.requests <- req
}

func (qs *QueueScheduler) GetWorker(worker chan handle.Request) {
	qs.workers <- worker
}

func (qs *QueueScheduler) Run() {
	qs.requests = make(chan handle.Request)
	qs.workers = make(chan chan handle.Request)

	go func() {
		var worker_queue []chan handle.Request
		var resquest_queue []handle.Request

		for {
			var free_worker chan handle.Request
			var free_request handle.Request

			if len(worker_queue) > 0 && len(resquest_queue) > 0 {
				free_request = resquest_queue[0]
				free_worker = worker_queue[0]
			}

			select {
			case req := <-qs.requests:
				resquest_queue = append(resquest_queue, req)
			case w := <-qs.workers:
				worker_queue = append(worker_queue, w)
			case free_worker <- free_request:
				resquest_queue = resquest_queue[1:]
				worker_queue = worker_queue[1:]
			}
		}
	}()
}
