package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"eaviwolph.com/StreamMusicDisplay/structs"
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
	errFunc := func() {
		log.Println("Default config creation")
		fileSavesConfBytes, err := json.Marshal(FileSavesConf)
		if err != nil {
			log.Printf("Error while marshalling FileSavesConf: %v", err)
			return
		}
		err = os.WriteFile("./saves/conf.json", fileSavesConfBytes, 0644)
		if err != nil {
			log.Printf("Error while writing conf.json: %v", err)
		}
	}

	b, err := os.ReadFile("./saves/conf.json")
	if err != nil {
		log.Printf("Error while reading conf.json: %v", err)
		errFunc()
		return
	}

	err = json.Unmarshal(b, &FileSavesConf)
	if err != nil {
		log.Printf("Error while unmarshalling conf.json: %v", err)
		errFunc()
		return
	}
}

func readTheme() {
	errFunc := func() {
		log.Println("Default theme creation")
		err := os.WriteFile("./saves/theme.txt", []byte(fmt.Sprintf("%d", Theme)), 0644)
		if err != nil {
			log.Printf("Error while writing theme.txt: %v", err)
		}
	}

	b, err := os.ReadFile("./saves/theme.txt")
	if err != nil {
		log.Printf("Error while reading theme.txt: %v", err)
		errFunc()
		return
	}

	Theme, err = strconv.Atoi(string(b))
	if err != nil {
		log.Printf("Error while parsing theme.txt: %v", err)
		errFunc()
		return
	}
}

func init() {
	readFileSavesConf()
	readTheme()
}
