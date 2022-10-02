package conf

import (
	"time"
)

type FileSaveStruct struct {
	Path    string
	Format  string
	Default string
}

type FileSaveConfStruct struct {
	SaveTxt   bool
	SaveImg   bool
	ImgPath   string
	FileSaves []FileSaveStruct
}

var (
	// Redirect URI for the Spotify API
	RedirectURI = "http://localhost:8888/callback"

	// Token config
	Code = ""

	// Config for file save
	// Global config
	Frequency = time.Second

	// File Save config
	FileSavesConf = FileSaveConfStruct{
		SaveTxt: true,
		SaveImg: true,
		ImgPath: "output/img.png",
		FileSaves: []FileSaveStruct{
			{
				Path:    "output/test.txt",
				Format:  "%artist% - %title% (%year%)",
				Default: "No song is currently playing",
			},
		},
	}
)
