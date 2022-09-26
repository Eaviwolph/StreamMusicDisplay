package serverHandler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func callbackHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/callback" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.URL.Query().Get("error") != "" {
		log.Default().Println("Code:", r.URL.Query().Get("error"))
		return
	}

	err := os.WriteFile("code", []byte(r.URL.Query().Get("code")), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Code written to file")
	fmt.Fprintf(w, "<a href=\"javascript:if (confirm(\"Close Window?\")) { close(); }\">close</a>")
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	path := "./static/index.html"
	if r.URL.Path != "/" {
		path = "./static" + r.URL.Path
		log.Println("Requested:", path)
	}
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Println("Requested error:", err)
	}

	w.Write(dat)
}

func StartServer() {
	http.HandleFunc("/callback", callbackHandle)
	http.HandleFunc("/", rootHandle)

	log.Println("Server listening on: http://localhost:8888")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
