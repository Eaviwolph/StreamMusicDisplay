package main

import (
	"context"
	"embed"
	"os"
	"time"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/requester"
	"eaviwolph.com/StreamMusicDisplay/serverHandler"
	"eaviwolph.com/StreamMusicDisplay/structs"
	"eaviwolph.com/StreamMusicDisplay/tools"
	"github.com/Scalingo/go-utils/logger"
)

//go:embed static/*
var staticFS embed.FS

func saveAllFiles(ctx context.Context, currentlyPlaying structs.CurrentlyPlaying) {
	if conf.FileSavesConf.SaveImg {
		ctx, log := logger.WithFieldToCtx(ctx, "ImgPath", conf.FileSavesConf.ImgPath)
		err := tools.SaveImgInFile(ctx, conf.FileSavesConf.ImgPath, currentlyPlaying)
		if err != nil {
			log.Printf("Error while saving %v: %v", conf.FileSavesConf.ImgPath, err)
		}
	}

	for _, config := range conf.FileSavesConf.FileSaves {
		ctx, log := logger.WithFieldToCtx(ctx, "TxtPath", config.Path)
		err := tools.SaveTxtInFile(ctx, config.Path, config.Format, config.Default, currentlyPlaying)
		if err != nil {
			log.Printf("Error while saving %v: %v", config.Path, err)
		}
	}
}

func main() {
	ctx := context.Background()
	log := logger.Get(ctx)

	go serverHandler.StartServer(ctx, staticFS)

	token := structs.AccessToken{}

	byteRefresh, err := os.ReadFile("./saves/refreshtoken")
	if err != nil {
		log.WithError(err).Error("Error while reading 'refreshtoken', requesting new one")
		requester.GetUserAuthorization(ctx)
	} else {
		token.RefreshToken = string(byteRefresh)
	}

	if token.RefreshToken == "" {
		for {
			if conf.Code != "" {
				time.Sleep(time.Duration(conf.FileSavesConf.Frequency))

				token, _ = requester.RequestAccessToken(ctx)

				if token != (structs.AccessToken{}) {
					os.WriteFile("./saves/refreshtoken", []byte(token.RefreshToken), 0644)
					break
				}
			}
		}
	}

	log.Info("Token received")

	for {
		time.Sleep(time.Duration(conf.FileSavesConf.Frequency))
		if conf.ExpireDate.Before(time.Now().Add(-1 * time.Minute)) {
			log.Info("Token expired, refreshing")
			token, _ = requester.RefreshAccessToken(ctx, token)
		}

		currentlyPlaying, err := requester.GetCurrentlyPlaying(ctx, token)
		if err != nil {
			continue
		}

		saveAllFiles(ctx, currentlyPlaying)
	}
}
