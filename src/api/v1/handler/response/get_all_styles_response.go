package response

type GetAllStylesResponse struct {
	BeerStyles []BeerStyle `json:"beer_styles"`
}

type BeerStyle struct {
	BeerStyle string `json:"beer_style"`
	Playlist  struct {
		Name   string `json:"name"`
		Tracks []struct {
			Name   string `json:"name"`
			Artist string `json:"artist"`
			Link   string `json:"link"`
		}
	}
}
