package serverHandler

import (
	"fmt"
	"log"
	"net/http"
)

var handleFunc = func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/callback" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	fmt.Println(r.URL.Path, ":", r.URL.Query())
}

func StartServer() {
	//requester.GetUserAuthorization()
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code

	http.HandleFunc("/callback", handleFunc)

	log.Println("Server listening on: http://localhost:8888")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
