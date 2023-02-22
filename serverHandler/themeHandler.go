package serverHandler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"github.com/Scalingo/go-utils/logger"
	"github.com/sirupsen/logrus"
)

func themeHandler(w http.ResponseWriter, r *http.Request) {
	_, log := logger.WithFieldsToCtx(r.Context(), logrus.Fields{
		"path":   r.URL.Path,
		"method": r.Method,
	})

	if r.Method == "GET" {
		fmt.Fprintf(w, "%d", conf.Theme)
	} else if r.Method == "POST" {
		if r.URL.Query().Get("num") == "" {
			log.Error("No num param")
			http.Error(w, "No num param", http.StatusBadRequest)
			return
		}

		var err error
		conf.Theme, err = strconv.Atoi(r.URL.Query().Get("num"))
		if err != nil {
			log.Errorf("Num is not an int: %v", err)
			http.Error(w, "Num is not an int", http.StatusBadRequest)
			return
		}

		err = os.WriteFile("./saves/theme.txt", []byte(fmt.Sprintf("%d", conf.Theme)), 0644)
		if err != nil {
			log.Errorf("Error while writing theme in file: %v", err)
			return
		}

		fmt.Fprintf(w, "Success")
	}
}
