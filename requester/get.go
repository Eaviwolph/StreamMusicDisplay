package requester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/structs"
	"eaviwolph.com/StreamMusicDisplay/tools"
)

func GetUserAuthorization() error {
	req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", nil)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	q := req.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", conf.ClientID)
	q.Add("scope", "user-read-currently-playing")
	q.Add("redirect_uri", conf.RedirectURI)

	req.URL.RawQuery = q.Encode()

	log.Printf("If you are not redirected automatically, please go to this link: %s", req.URL.String())

	tools.OpenBrowser(req.URL.String())
	return nil
}

func GetCurrentlyPlaying(token structs.AccessToken) (structs.CurrentlyPlaying, error) {
	currentlyPlaying := structs.CurrentlyPlaying{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/currently-playing", nil)
	if err != nil {
		log.Printf("fail to create request: %v", err)
		return currentlyPlaying, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token.AccessToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("fail to send request: %v", err)
		return currentlyPlaying, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Printf("fail to read response body: %v", err)
		return currentlyPlaying, err
	}

	err = json.Unmarshal(body, &currentlyPlaying)
	if err != nil && currentlyPlaying.IsPlaying == false {
		currentlyPlaying.Item = structs.Item{Name: "Nothing is playing"}
	}

	return currentlyPlaying, err
}
