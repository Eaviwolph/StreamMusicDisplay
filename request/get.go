package requester

import (
	"fmt"
	"log"
	"net/http"

	"eaviwolph.com/StreamMusicDisplay/conf"
	"eaviwolph.com/StreamMusicDisplay/tools"
)

func GetUserAuthorization() error {
	req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", nil)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	q := req.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", conf.ClientID)
	q.Add("scope", "user-read-currently-playing")
	q.Add("redirect_uri", conf.RedirectURI)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	tools.OpenBrowser(req.URL.String())
	return nil
}
