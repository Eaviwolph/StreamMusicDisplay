package tools

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"strings"

	"eaviwolph.com/StreamMusicDisplay/structs"
	"github.com/Scalingo/go-utils/errors"
)

func SaveImgInFile(ctx context.Context, path string, cur structs.CurrentlyPlaying) error {
	if len(cur.Item.Album.Images) == 0 {
		return errors.Wrapf(ctx, nil, "currentlyplaying has no images")
	}

	URI := strings.Split(path, "/")
	folderPath := strings.Join(URI[:len(URI)-1], "/")
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to create folder %s", folderPath)
	}

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to create file %s", path)
	}

	defer f.Close()

	httpClient := http.Client{}
	resp, err := httpClient.Get(cur.Item.Album.Images[0].URL)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to get image from '%s'", cur.Item.Album.Images[0].URL)
	}

	defer resp.Body.Close()

	_, err = f.ReadFrom(resp.Body)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to read image from '%s'", cur.Item.Album.Images[0].URL)
	}
	return nil
}

func SaveTxtDefaultInFile(ctx context.Context, path string, def string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to create file %s", path)
	}

	defer f.Close()

	_, err = f.WriteString(def)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to write default text in file '%s'", path)
	}

	return nil
}

func SaveTxtInFile(ctx context.Context, path string, format string, def string, cur structs.CurrentlyPlaying) error {
	URI := strings.Split(path, "/")
	folderPath := strings.Join(URI[:len(URI)-1], "/")
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to create folder '%s'", folderPath)
	}

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(ctx, err, "fail to create file '%s'", path)
	}

	if format == "" {
		return nil
	}

	if cur.Item.Name == "" {
		_, err = f.WriteString(def)
		if err != nil {
			return errors.Wrapf(ctx, err, "fail to write default text in file '%s'", path)
		}
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
		return errors.Wrapf(ctx, err, "fail to write text in file '%s'", path)
	}
	return nil
}
