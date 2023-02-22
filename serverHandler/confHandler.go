package serverHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/structs"
	"github.com/Scalingo/go-utils/logger"
	"github.com/sirupsen/logrus"
)

func confHandler(w http.ResponseWriter, r *http.Request) {
	_, log := logger.WithFieldsToCtx(r.Context(), logrus.Fields{
		"path":   r.URL.Path,
		"method": r.Method,
	})

	if r.Method == "POST" {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Error("Error while reading body")
			http.Error(w, "Error while reading body", http.StatusInternalServerError)
			return
		}

		log.Debug("Conf written to file")
		parsed := structs.FileSaveConfStruct{}
		err = json.Unmarshal(bodyBytes, &parsed)
		if err != nil {
			log.WithError(err).Error("Error while parsing body")
			http.Error(w, "Error while parsing body", http.StatusInternalServerError)
			return
		}

		conf.FileSavesConf = parsed
		err = os.WriteFile("./saves/conf.json", bodyBytes, 0644)
		if err != nil {
			log.WithError(err).Error("Error while writing conf in file")
			http.Error(w, "Error while writing conf in file", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Success")
	} else if r.Method == "GET" {
		b, err := json.Marshal(conf.FileSavesConf)
		if err != nil {
			log.WithError(err).Error("Error while parsing conf")
			http.Error(w, "Error while parsing conf", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	}
}
