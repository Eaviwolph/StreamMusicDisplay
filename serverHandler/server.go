package serverHandler

import (
	"context"
	"embed"
	"fmt"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/sirupsen/logrus"
)

var staticFS embed.FS

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, log := logger.WithFieldsToCtx(r.Context(), logrus.Fields{
		"path":   r.URL.Path,
		"method": r.Method,
	})

	path := "static/index.html"
	if r.URL.Path != "/" {
		path = "static" + r.URL.Path
		log.Infof("Requested: %v", path)
	}
	dat, err := staticFS.ReadFile(path)
	if err != nil {
		log.WithError(err).Error("Error while reading")
		http.Error(w, fmt.Sprintf("Error while reading %v", path), http.StatusInternalServerError)
	}

	w.Write(dat)
}

func StartServer(ctx context.Context, static embed.FS) {
	log := logger.Get(ctx)
	staticFS = static
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/conf", confHandler)
	http.HandleFunc("/theme", themeHandler)

	http.HandleFunc("/", rootHandler)

	log.Info("Server listening on: http://localhost:8888")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
