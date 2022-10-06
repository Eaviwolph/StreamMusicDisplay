package conf

import (
	"encoding/json"
	"log"
	"os"
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
		Frequency: time.Second,
		SaveTxt:   true,
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
)

func init() {
	defer func() {
		if r := recover(); r == nil {
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
	}()

	b, err := os.ReadFile("./saves/conf.json")
	if err != nil {
		log.Printf("Error while reading conf.json: %v", err)
		return
	}

	err = json.Unmarshal(b, &FileSavesConf)
	if err != nil {
		log.Printf("Error while unmarshalling conf.json: %v", err)
		return
	}
}
