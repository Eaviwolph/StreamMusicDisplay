package tools

import (
	"fmt"
	"net/http"
	"os"

	"strings"

	"github.com/zmb3/spotify/v2"
)

func SaveImgInFile(path string, cur *spotify.CurrentlyPlaying) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	httpClient := http.Client{}
	resp, err := httpClient.Get(cur.Item.Album.Images[0].URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = f.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func SaveTxtInFile(path string, format string, cur *spotify.CurrentlyPlaying) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	format = strings.ReplaceAll(format, "%artist%", cur.Item.Artists[0].Name)
	format = strings.ReplaceAll(format, "%title%", cur.Item.Name)
	format = strings.ReplaceAll(format, "%album%", cur.Item.Album.Name)
	format = strings.ReplaceAll(format, "%year%", cur.Item.Album.ReleaseDate)
	format = strings.ReplaceAll(format, "%track%", fmt.Sprintf("%d", cur.Item.TrackNumber))
	format = strings.ReplaceAll(format, "%duration%", fmt.Sprintf("%d", cur.Item.Duration))
	format = strings.ReplaceAll(format, "%progress%", fmt.Sprintf("%d", cur.Progress))

	_, err = f.WriteString(format)
	if err != nil {
		return err
	}
	return nil
}
