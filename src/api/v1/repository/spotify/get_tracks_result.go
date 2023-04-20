package repository

type GetTracksResult struct {
	Items []struct {
		Track struct {
			Artists      Artists `json:"artists"`
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

func (a Artist) ToString() string {
	return a.Name
}

type Artists []Artist

func (a Artists) ToString() string {
	var result string
	for _, artist := range a {
		result += artist.ToString() + ", "
	}
	return result[:len(result)-2]
}
