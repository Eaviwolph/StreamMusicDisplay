package main

import (
	"sync"

	"eaviwolph.com/StreamMusicDisplay/requester"
	"eaviwolph.com/StreamMusicDisplay/serverHandler"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go serverHandler.StartServer()

	requester.GetUserAuthorization()
	wg.Wait()
}
