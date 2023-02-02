package schedulers

import "reptile/handle"

type NormalScheduler struct {
	requests chan handle.Request
}

func (ns *NormalScheduler) Submit(req handle.Request) {
	//ns.requests <- req 出现死锁
	go func() {
		ns.requests <- req
	}()
}

func (ns *NormalScheduler) GetChan(requests chan handle.Request) {
	ns.requests = requests
}
