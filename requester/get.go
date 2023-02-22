package requester

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/structs"
	"eaviwolph.com/StreamMusicDisplay/tools"
	"github.com/Scalingo/go-utils/logger"
)

func GetUserAuthorization(ctx context.Context) error {
	log := logger.Get(ctx)

	req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", nil)
	if err != nil {
		log.WithError(err).Error("fail to create request")
		return err
	}

	q := req.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", conf.ClientID)
	q.Add("scope", "user-read-currently-playing")
	q.Add("redirect_uri", conf.RedirectURI)

	req.URL.RawQuery = q.Encode()

	log.Infof("If you are not redirected automatically, please go to this link: %s", req.URL.String())

	tools.OpenBrowser(ctx, req.URL.String())
	return nil
}

func GetCurrentlyPlaying(ctx context.Context, token structs.AccessToken) (structs.CurrentlyPlaying, error) {
	log := logger.Get(ctx)

	currentlyPlaying := structs.CurrentlyPlaying{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/currently-playing", nil)
	if err != nil {
		log.WithError(err).Error("fail to create request")
		return currentlyPlaying, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token.AccessToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("fail to send request")
		return currentlyPlaying, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.WithError(err).Error("fail to read response body")
		return currentlyPlaying, err
	}

	err = json.Unmarshal(body, &currentlyPlaying)
	if err != nil && !currentlyPlaying.IsPlaying {
		currentlyPlaying.Item = structs.Item{Name: ""}
	}

	return currentlyPlaying, err
}
