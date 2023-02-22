package serverHandler

import (
	"net/http"
	"os"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"github.com/Scalingo/go-utils/logger"
	"github.com/sirupsen/logrus"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	_, log := logger.WithFieldsToCtx(r.Context(), logrus.Fields{
		"path":   r.URL.Path,
		"method": r.Method,
	})

	if r.URL.Query().Get("error") != "" {
		log.Errorf("Code: %v", r.URL.Query().Get("error"))
		http.Error(w, "Error while getting code", http.StatusBadRequest)
		return
	}

	conf.Code = r.URL.Query().Get("code")

	err := os.WriteFile("saves/code", []byte(r.URL.Query().Get("code")), 0644)
	if err != nil {
		log.WithError(err).Error("fail to write code in file")
		http.Error(w, "Error while writing code in file", http.StatusInternalServerError)
		return
	}
	log.Info("Code written to file")

	dat, err := staticFS.ReadFile("static/callback.html")
	if err != nil {
		log.WithError(err).Error("Requested error")
		http.Error(w, "Error while reading callback.html", http.StatusInternalServerError)
		return
	}
	w.Write(dat)
}
