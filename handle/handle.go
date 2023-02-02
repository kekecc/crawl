package handle

import (
	"fmt"
	"log"
	"reptile/fetch"
)

type Handler struct { //low
}

func (h *Handler) Run(urls ...Request) {
	var resquests []Request
	// for _, res := range urls {
	// 	resquests = append(resquests, res)
	// }
	resquests = append(resquests, urls...)

	for len(resquests) > 0 {
		res := resquests[0]
		resquests = resquests[1:]

		log.Println("fetch url: ", res.Url)
		body, err := fetch.Fetch(res.Url)
		if err != nil {
			log.Println("fetch err: ", err)
			return
		}
		result := res.ParseFunc(body) //会有其他任务
		resquests = append(resquests, result.Requests...)

		for _, content := range result.Contents {
			fmt.Println("get:", string(content.([]byte)))
		}
	}
}
