package main

import (
	"log"
	"time"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/requester"
	"eaviwolph.com/StreamMusicDisplay/serverHandler"
	"eaviwolph.com/StreamMusicDisplay/structs"
	"eaviwolph.com/StreamMusicDisplay/tools"
)

func saveAllFiles(currentlyPlaying structs.CurrentlyPlaying) error {
	err := tools.SaveImgInFile(conf.FileSavesConf.ImgPath, currentlyPlaying)
	if err != nil {
		log.Printf("Error while saving %v: %v", conf.FileSavesConf.ImgPath, err)
		return err
	}
	for _, config := range conf.FileSavesConf.FileSaves {
		err := tools.SaveTxtInFile(config.Path, config.Format, currentlyPlaying)
		if err != nil {
			log.Printf("Error while saving %v: %v", config.Path, err)
			return err
		}
	}
	return nil
}

func main() {
	go serverHandler.StartServer()

	requester.GetUserAuthorization()

	token := structs.AccessToken{}

	for {
		if conf.Code != "" {
			time.Sleep(conf.FileSavesConf.Frequency)

			token, _ = requester.RequestAccessToken()
			if token != (structs.AccessToken{}) {
				break
			}
		}
	}

	for {
		time.Sleep(conf.FileSavesConf.Frequency)
		if conf.ExpireDate.Before(time.Now()) {
			token, _ = requester.RefreshAccessToken(token)
		}

		currentlyPlaying, err := requester.GetCurrentlyPlaying(token)
		if err != nil {
			continue
		}

		err = saveAllFiles(currentlyPlaying)
		if err != nil {
			continue
		}
	}
}
