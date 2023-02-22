package conf

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"eaviwolph.com/StreamMusicDisplay/structs"
	"github.com/Scalingo/go-utils/logger"
)

var (
	// Redirect URI for the Spotify API
	RedirectURI = "http://localhost:8888/callback"

	// Token config
	Code = ""

	// Config for file save
	// Global config

	ExpireDate = time.Time{}

	// File Save config
	FileSavesConf = structs.FileSaveConfStruct{
		Frequency: int(time.Second),
		SaveImg:   true,
		ImgPath:   "output/img.png",
		FileSaves: []structs.FileSaveStruct{
			{
				Path:    "output/test.txt",
				Format:  "%artist% - %title% (%year%)",
				Default: "No song is currently playing",
			},
		},
	}

	Theme = 0
)

func readFileSavesConf() {
	ctx := context.Background()
	log := logger.Get(ctx)
	errFunc := func() {
		log.Info("Default config creation")
		fileSavesConfBytes, err := json.Marshal(FileSavesConf)
		if err != nil {
			log.WithError(err).Error("Error while marshalling FileSavesConf")
			return
		}
		err = os.WriteFile("./saves/conf.json", fileSavesConfBytes, 0644)
		if err != nil {
			log.WithError(err).Error("Error while writing conf.json")
		}
	}

	b, err := os.ReadFile("./saves/conf.json")
	if err != nil {
		log.WithError(err).Error("Error while reading conf.json")
		errFunc()
		return
	}

	err = json.Unmarshal(b, &FileSavesConf)
	if err != nil {
		log.WithError(err).Error("Error while unmarshalling conf.json")
		errFunc()
		return
	}
}

func readTheme() {
	ctx := context.Background()
	log := logger.Get(ctx)

	errFunc := func() {
		log.Println("Default theme creation")
		err := os.WriteFile("./saves/theme.txt", []byte(fmt.Sprintf("%d", Theme)), 0644)
		if err != nil {
			log.WithError(err).Error("Error while writing theme.txt")
		}
	}

	b, err := os.ReadFile("./saves/theme.txt")
	if err != nil {
		log.WithError(err).Error("Error while reading theme.txt")
		errFunc()
		return
	}

	Theme, err = strconv.Atoi(string(b))
	if err != nil {
		log.WithError(err).Error("Error while parsing theme.txt")
		errFunc()
		return
	}
}

func init() {
	os.Mkdir("./saves", 0755)
	readFileSavesConf()
	readTheme()
}
