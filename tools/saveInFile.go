package tools

import (
	"fmt"
	"net/http"
	"os"

	"strings"

	"eaviwolph.com/StreamMusicDisplay/structs"
)

func SaveImgInFile(path string, cur structs.CurrentlyPlaying) error {
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

func SaveTxtDefaultInFile(path string, def string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(def)
	if err != nil {
		return err
	}

	return nil
}

func SaveTxtInFile(path string, format string, cur structs.CurrentlyPlaying) error {
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
	format = strings.ReplaceAll(format, "%duration%", fmt.Sprintf("%d", cur.Item.DurationMs/1000))
	format = strings.ReplaceAll(format, "%progress%", fmt.Sprintf("%d", cur.ProgressMs/1000))

	_, err = f.WriteString(format)
	if err != nil {
		return err
	}
	return nil
}
