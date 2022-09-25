package requester

import (
	"fmt"
	"log"
	"net/http"
)

func GetUserAuthorization() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", client_id)
	q.Add("scope", "user-read-private user-read-email")
	q.Add("redirect_uri", "http://localhost:8888/callback")

	req.URL.RawQuery = q.Encode()

	client.Do(req)

	fmt.Println(req.URL.String())
}
