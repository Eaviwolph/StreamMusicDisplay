package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/serverHandler"
	"eaviwolph.com/StreamMusicDisplay/tools"
	"github.com/zmb3/spotify/v2"
)

func RefreshToken(ctx context.Context, client *spotify.Client) (*spotify.Client, error) {
	if conf.Expiry.Before(time.Now()) {
		r, _ := http.NewRequest("", "", nil)
		q := r.URL.Query()
		q.Add("code", conf.Code)
		r.URL.RawQuery = q.Encode()
		tok, err := conf.Auth.Token(ctx, conf.State, r)
		if err != nil {
			log.Printf("fail to refresh token: %v\n", err)
			return client, err
		}
		client = spotify.New(conf.Auth.Client(ctx, tok))
	}
	return client, nil
}

func SaveInFile(cur *spotify.CurrentlyPlaying, currentlyPlayingID string) (string, error) {
	if conf.FileSavesConf.SaveTxt {
		for _, fs := range conf.FileSavesConf.FileSaves {
			err := tools.SaveTxtInFile(fs.Path, fs.Format, cur)
			if err != nil {
				log.Printf("fail to save in file: %v", err)
				return currentlyPlayingID, err
			}
		}
	}

	if conf.FileSavesConf.SaveImg && currentlyPlayingID != cur.Item.ID.String() {
		err := tools.SaveImgInFile(conf.FileSavesConf.ImgPath, cur)
		if err != nil {
			log.Printf("fail to save in file: %v", err)
			return currentlyPlayingID, err
		}
		currentlyPlayingID = cur.Item.ID.String()
	}
	return currentlyPlayingID, nil
}

func main() {
	ctx := context.Background()
	os.Setenv("SPOTIFY_ID", "4e8e40f5b0b7455b97ffb08b0a5fd347")
	os.Setenv("SPOTIFY_SECRET", "3ba9c6b6662541ad84db0871e9cc6f09")

	go serverHandler.StartServer()

	url := conf.Auth.AuthURL(conf.State)

	tools.OpenBrowser(url)

	// wait for auth to complete
	client := <-conf.Ch

	currentlyPlayingID := ""
	for {
		time.Sleep(conf.Frequency)

		client, err := RefreshToken(ctx, client)
		if err != nil {
			continue
		}

		cur, err := client.PlayerCurrentlyPlaying(ctx)
		if err != nil || cur == nil || cur.Item == nil {
			log.Printf("fail to get current playing: %v", err)
			continue
		}

		currentlyPlayingID, _ = SaveInFile(cur, currentlyPlayingID)
	}
}
