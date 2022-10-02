package main

import (
	"context"
	"log"
	"os"
	"time"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/serverHandler"
	"eaviwolph.com/StreamMusicDisplay/tools"
	"github.com/zmb3/spotify/v2"
)

func RefreshToken(ctx context.Context) (*spotify.Client, error) {
	log.Printf("refreshing token")
	conf.Token, err := conf.Auth.Token(context.Background(), conf.State, r)
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

func SaveDefaultInFile() error {
	for _, fs := range conf.FileSavesConf.FileSaves {
		err := tools.SaveTxtDefaultInFile(fs.Path, fs.Default)
		if err != nil {
			log.Printf("fail to save in file: %v", err)
			return err
		}
	}
	return nil
}

func main() {
	ctx := context.Background()

	err := os.Setenv("SPOTIFY_ID", conf.ClientID)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("SPOTIFY_SECRET", conf.ClientSecret)
	if err != nil {
		log.Fatal(err)
	}

	conf.InitAuth()

	go serverHandler.StartServer()

	url := conf.Auth.AuthURL(conf.State)
	tools.OpenBrowser(url)

	// wait for auth to complete
	client := <-conf.Ch

	currentlyPlayingID := ""
	for {
		time.Sleep(conf.Frequency)

		if conf.Token.Expiry.Before(time.Now()) {
			client, err = RefreshToken(ctx)
			if err != nil {
				continue
			}
		}

		cur, err := client.PlayerCurrentlyPlaying(ctx)
		if err != nil {
			log.Printf("fail to get current playing: %v", err)
			continue
		}
		if cur == nil || cur.Item == nil {
			SaveDefaultInFile()
			continue
		}

		currentlyPlayingID, _ = SaveInFile(cur, currentlyPlayingID)
	}
}
