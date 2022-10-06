package structs

type Context struct {
	Type string `json:"type"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Album struct {
	AlbumType   string  `json:"album_type"`
	TotalTracks int     `json:"total_tracks"`
	Images      []Image `json:"images"`
	Name        string  `json:"name"`
	Popularity  int     `json:"popularity"`
	ReleaseDate string  `json:"release_date"`
}

type Artist struct {
	Genre      string  `json:"genre"`
	Images     []Image `json:"images"`
	Name       string  `json:"name"`
	Popularity int     `json:"popularity"`
}

type Item struct {
	Album       Album    `json:"album"`
	Artists     []Artist `json:"artists"`
	DiscNumber  int      `json:"disc_number"`
	DurationMs  int      `json:"duration_ms"`
	Explicit    bool     `json:"explicit"`
	Name        string   `json:"name"`
	Popularity  int      `json:"popularity"`
	TrackNumber int      `json:"track_number"`
}

type CurrentlyPlaying struct {
	Repeat       string `json:"repeat_state"`
	ShuffleState string `json:"shuffle_state"`
	Timestamp    int    `json:"timestamp"`
	ProgressMs   int    `json:"progress_ms"`
	IsPlaying    bool   `json:"is_playing"`
	Item         Item   `json:"item"`
}
