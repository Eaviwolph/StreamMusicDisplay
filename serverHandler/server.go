package serverHandler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"github.com/zmb3/spotify/v2"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	tok, err := conf.Auth.Token(r.Context(), conf.State, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	conf.Code = r.URL.Query().Get("code")
	conf.Expiry = tok.Expiry

	if st := r.FormValue("state"); st != conf.State {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, conf.State)
	}

	// use the token to get an authenticated client
	client := spotify.New(conf.Auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	conf.Ch <- client
}

func saveConfHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("saveConfHandler")
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
	http.HandleFunc("/saveConf", saveConfHandler)

	http.HandleFunc("/", rootHandler)

	log.Println("Server listening on: http://localhost:8888")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
