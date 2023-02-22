package main

import (
	"context"
	"embed"
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

	go serverHandler.StartServer(ctx, staticFS)

	requester.GetUserAuthorization(ctx)

	token := structs.AccessToken{}

	for {
		if conf.Code != "" {
			time.Sleep(time.Duration(conf.FileSavesConf.Frequency))

			token, _ = requester.RequestAccessToken(ctx)
			if token != (structs.AccessToken{}) {
				break
			}
		}
	}

	for {
		time.Sleep(time.Duration(conf.FileSavesConf.Frequency))
		if conf.ExpireDate.Before(time.Now().Add(-1 * time.Minute)) {
			token, _ = requester.RefreshAccessToken(ctx, token)
		}

		currentlyPlaying, err := requester.GetCurrentlyPlaying(ctx, token)
		if err != nil {
			continue
		}

		saveAllFiles(ctx, currentlyPlaying)
	}
}
