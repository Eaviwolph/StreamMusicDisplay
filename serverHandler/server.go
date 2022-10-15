package serverHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/structs"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("error") != "" {
		log.Println("Code:", r.URL.Query().Get("error"))
		http.Error(w, "Error while getting code", http.StatusBadRequest)
		return
	}

	conf.Code = r.URL.Query().Get("code")
	err := os.WriteFile("code", []byte(r.URL.Query().Get("code")), 0644)
	if err != nil {
		log.Printf("fail to write code in file: %v", err)
		http.Error(w, "Error while writing code in file", http.StatusInternalServerError)
		return
	}
	log.Println("Code written to file")

	dat, err := os.ReadFile("./static/callback.html")
	if err != nil {
		log.Println("Requested error:", err)
		http.Error(w, "Error while reading callback.html", http.StatusInternalServerError)
		return
	}
	w.Write(dat)
}

func confHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error while reading body:", err)
			http.Error(w, "Error while reading body", http.StatusInternalServerError)
			return
		}

		log.Println("Conf written to file")
		parsed := structs.FileSaveConfStruct{}
		err = json.Unmarshal(bodyBytes, &parsed)
		if err != nil {
			log.Println("Error while parsing body:", err)
			http.Error(w, "Error while parsing body", http.StatusInternalServerError)
			return
		}

		conf.FileSavesConf = parsed
		err = os.WriteFile("./saves/conf.json", bodyBytes, 0644)
		if err != nil {
			log.Println("Error while writing conf in file:", err)
			http.Error(w, "Error while writing conf in file", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Success")
	}

	if r.Method == "GET" {
		b, err := json.Marshal(conf.FileSavesConf)
		if err != nil {
			log.Println("Error while parsing conf:", err)
			http.Error(w, "Error while parsing conf", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	}
}

func themeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, fmt.Sprintf("%d", conf.Theme))
	} else if r.Method == "POST" {
		if r.URL.Query().Get("num") == "" {
			http.Error(w, "No num param", http.StatusBadRequest)
			return
		}
		var err error
		conf.Theme, err = strconv.Atoi(r.URL.Query().Get("num"))
		if err != nil {
			http.Error(w, "Num is not an int", http.StatusBadRequest)
			return
		}

		err = os.WriteFile("./saves/theme.txt", []byte(fmt.Sprintf("%d", conf.Theme)), 0644)
		if err != nil {
			log.Printf("Error while writing theme.txt: %v", err)
		}

		fmt.Fprintf(w, "Success")
	}
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
		http.Error(w, fmt.Sprintf("Error while reading %v", path), http.StatusInternalServerError)
	}

	w.Write(dat)
}

func StartServer() {
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/conf", confHandler)
	http.HandleFunc("/theme", themeHandler)

	http.HandleFunc("/", rootHandler)

	log.Println("Server listening on: http://localhost:8888")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
