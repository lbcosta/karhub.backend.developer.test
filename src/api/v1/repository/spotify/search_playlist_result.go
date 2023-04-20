package repository

type SearchPlaylistsResult struct {
	Playlists struct {
		Items []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"playlists"`
}
