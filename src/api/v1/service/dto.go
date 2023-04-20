package service

type GetTokenResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type SearchPlaylistsResult struct {
	Playlists struct {
		Items []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"playlists"`
}

type GetTracksResult struct {
	Items []struct {
		Track struct {
			Artists      []Artist `json:"artists"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
}

type Artist struct {
	Name string `json:"name"`
}

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}

type Playlist struct {
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

type BeerPlaylist struct {
	BeerStyle string   `json:"beer_style"`
	Playlist  Playlist `json:"playlist"`
}
