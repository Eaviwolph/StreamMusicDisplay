package requester

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/structs"
)

func RequestAccessToken() (structs.AccessToken, error) {
	token := structs.AccessToken{}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", conf.Code)
	data.Set("redirect_uri", conf.RedirectURI)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("fail to create request: %v", err)
		return token, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth(conf.ClientID, conf.ClientSecret)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("fail to send request: %v", err)
		return token, err
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	resp.Body.Close()
	if err != nil {
		log.Printf("fail to read response body: %v", err)
		return token, err
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Printf("fail to unmarshal response body: %v", err)
		return token, err
	}
	conf.ExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	return token, nil
}

func RefreshAccessToken(token structs.AccessToken) (structs.AccessToken, error) {
	newToken := structs.AccessToken{}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("code", conf.Code)
	data.Set("refresh_token", token.RefreshToken)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("fail to create request: %v", err)
		return token, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth(conf.ClientID, conf.ClientSecret)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("fail to send request: %v", err)
		return token, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Printf("fail to read response body: %v", err)
		return token, err
	}

	err = json.Unmarshal(body, &newToken)
	if err != nil {
		log.Printf("fail to unmarshal response body: %v", err)
		return newToken, err
	}
	token.AccessToken = newToken.AccessToken
	token.ExpiresIn = newToken.ExpiresIn
	conf.ExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	return token, nil
}
