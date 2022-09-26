package requester

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RequestAccessToken() error {
	code, err := os.ReadFile("./code")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	var jsonStr = []byte(`{"grant_type":"authorization_code","code":"` + string(code) + `"}`)
	resp, err := http.Post("https://accounts.spotify.com/api/token", "application/x-www-form-urlencoded", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	fmt.Println(resp.Body)
	return nil
}
