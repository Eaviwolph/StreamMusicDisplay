package tools

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/Scalingo/go-utils/logger"
)

func OpenBrowser(ctx context.Context, url string) {
	var err error
	log := logger.Get(ctx)

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.WithError(err).Info("Error while opening browser")
	}
}
