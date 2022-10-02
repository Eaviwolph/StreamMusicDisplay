package requester

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"eaviwolph.com/StreamMusicDisplay/conf"
)

func RequestAccessToken() error {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", conf.Code)
	data.Set("redirect_uri", conf.RedirectURI)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(conf.ClientID+":"+conf.ClientSecret)))

	var ch chan<- interface{}
	ch <- client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
	return nil
}
