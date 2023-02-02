package datasave

import "log"

func Save() chan interface{} {
	saves := make(chan interface{})

	go func() {
		for {
			mes := <-saves
			log.Println("get: ", mes.(string))
		}
	}()

	return saves
}
