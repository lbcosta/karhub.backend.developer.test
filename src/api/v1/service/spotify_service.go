package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BeersPlaylists map[string]SearchPlaylistsResult

type SpotifyService struct{}

func NewSpotifyService() SpotifyService {
	return SpotifyService{}
}

func (s SpotifyService) GetToken() (string, error) {
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	data.Add("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))

	encodedData := data.Encode()

	authUrl := "https://accounts.spotify.com/api/token"
	req, err := http.NewRequest("POST", authUrl, strings.NewReader(encodedData))
	if err != nil {
		return "", err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result GetTokenResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return result.AccessToken, nil
}

func (s SpotifyService) SearchPlaylists(beerStyles []string, token string) ([]BeerPlaylist, error) {
	playlists := make([]BeerPlaylist, len(beerStyles))

	for idx, beerStyle := range beerStyles {
		searchResult, err := s.searchPlaylist(beerStyle, token)
		if err != nil {
			return nil, err
		}

		playListId := searchResult.Playlists.Items[0].Id
		tracks, err := s.getTracks(playListId, token)
		if err != nil {
			return nil, err
		}

		playlists[idx] = BeerPlaylist{
			BeerStyle: beerStyle,
			Playlist: Playlist{
				Name:   searchResult.Playlists.Items[0].Name,
				Tracks: tracks,
			},
		}
	}

	return playlists, nil
}

func (s SpotifyService) searchPlaylist(query, token string) (SearchPlaylistsResult, error) {
	searchUrl := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=playlist&limit=1", query)
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return SearchPlaylistsResult{}, err
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", token)},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SearchPlaylistsResult{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return SearchPlaylistsResult{}, fmt.Errorf("failed to search playlist: %s", resp.Status)
	}

	var searchResult SearchPlaylistsResult
	err = json.NewDecoder(resp.Body).Decode(&searchResult)
	if err != nil {
		return SearchPlaylistsResult{}, err
	}

	return searchResult, nil
}

func (s SpotifyService) getTracks(playlistId, token string) ([]Track, error) {
	searchUrl := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistId)
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", token)},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get tracks: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchResult GetTracksResult
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return nil, err
	}

	var tracks = make([]Track, len(searchResult.Items))
	for idx, track := range searchResult.Items {
		artists := concatArtists(track.Track.Artists)

		tracks[idx] = Track{
			Name:   track.Track.Name,
			Artist: artists,
			Link:   track.Track.ExternalUrls.Spotify,
		}
	}

	return tracks, nil
}

func concatArtists(artists []Artist) string {
	var str []string
	for _, artist := range artists {
		str = append(str, artist.Name)
	}

	return strings.Join(str, ", ")
}
