package conf

import (
	"time"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type FileSaveStruct struct {
	Path   string
	Format string
}

type FileSaveConfStruct struct {
	SaveTxt   bool
	SaveImg   bool
	ImgPath   string
	FileSaves []FileSaveStruct
}

var (
	// Redirect URI for the Spotify API
	redirectURI = "http://localhost:8888/callback"

	// Auth is the Spotify authentication object
	Auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying),
	)

	// Ch is the channel to send the authenticated client
	Ch = make(chan *spotify.Client)

	// State is the state for the Spotify authentication
	State = "abc123"

	// Token config
	Code   = ""
	Expiry = time.Time{}

	// Global config
	Frequency = time.Second

	// File Save config
	FileSavesConf = FileSaveConfStruct{
		SaveTxt: true,
		SaveImg: true,
		ImgPath: "output/img.png",
		FileSaves: []FileSaveStruct{
			{
				Path:   "output/test.txt",
				Format: "%artist% - %title% (%year%)",
			},
		},
	}
)
