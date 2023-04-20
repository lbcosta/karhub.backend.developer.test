package domain

type Playlist struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}
