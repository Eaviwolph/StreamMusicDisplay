package requester

import (
	"fmt"
	"io"
	"io/ioutil"
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
	data.Set("client_id", conf.ClientID)
	data.Set("client_secret", conf.ClientSecret)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth(conf.ClientID, conf.ClientSecret)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(body))
	return nil
}
