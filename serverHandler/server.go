package serverHandler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"eaviwolph.com/StreamMusicDisplay/conf"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-------------------callbackHandler")
	if r.URL.Path != "/callback" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.URL.Query().Get("error") != "" {
		log.Default().Println("Code:", r.URL.Query().Get("error"))
		return
	}

	conf.Code = r.URL.Query().Get("code")
	err := os.WriteFile("code", []byte(r.URL.Query().Get("code")), 0644)
	if err != nil {
		log.Printf("fail to write code in file: %v", err)
		return
	}
	log.Println("Code written to file")

	dat, err := os.ReadFile("./static/callback.html")
	if err != nil {
		log.Println("Requested error:", err)
	}
	w.Write(dat)
}

func confHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("saveConfHandler")
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("refreshTokenHandler")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/saveConf", confHandler)
	http.HandleFunc("/refreshToken", refreshTokenHandler)

	http.HandleFunc("/", rootHandler)

	log.Println("Server listening on: http://localhost:8888")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
