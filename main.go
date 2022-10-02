package main

import (
	"time"

	"eaviwolph.com/StreamMusicDisplay/conf"
	requester "eaviwolph.com/StreamMusicDisplay/request"
	"eaviwolph.com/StreamMusicDisplay/serverHandler"
)

func main() {
	go serverHandler.StartServer()

	requester.GetUserAuthorization()

	for {
		if conf.Code != "" {
			time.Sleep(conf.Frequency)

			requester.RequestAccessToken()
		}
	}
}
