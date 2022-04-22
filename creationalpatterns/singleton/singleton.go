package singleton

import "time"

func Run() {

	for i := 0; i < 5; i++ {
		go getInstance()
	}

	time.Sleep(time.Second * 5)
}
