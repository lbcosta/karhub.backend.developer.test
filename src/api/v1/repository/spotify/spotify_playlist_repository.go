package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"karhub.backend.developer.test/src/api/v1/model"
	"karhub.backend.developer.test/src/api/v1/repository"
	"net/http"
	"strings"
)

const (
	getPlaylistUrl = "https://api.spotify.com/v1/search?q=%s&type=playlist&limit=1"
	getTracksUrl   = "https://api.spotify.com/v1/playlists/%s/tracks"
)

type SpotifyPlaylistRepository struct{}

func NewSpotifyPlaylistRepository() repository.PlaylistRepository {
	return SpotifyPlaylistRepository{}
}

func (SpotifyPlaylistRepository) GetPlaylist(query, token string) (model.Playlist, error) {
	url := fmt.Sprintf(getPlaylistUrl, strings.ReplaceAll(query, " ", "+"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Playlist{}, err
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", token)},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.Playlist{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return model.Playlist{}, fmt.Errorf("failed to search playlist: %s", resp.Status)
	}

	var searchResult SearchPlaylistsResult
	err = json.NewDecoder(resp.Body).Decode(&searchResult)
	if err != nil {
		return model.Playlist{}, err
	}

	playlist := model.Playlist{
		ID:     searchResult.Playlists.Items[0].Id,
		Name:   searchResult.Playlists.Items[0].Name,
		Tracks: nil,
	}

	return playlist, nil
}

func (SpotifyPlaylistRepository) GetTracks(playlistId, token string) (model.Tracks, error) {
	url := fmt.Sprintf(getTracksUrl, playlistId)

	req, err := http.NewRequest("GET", url, nil)
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

	var getTracksResult GetTracksResult
	err = json.Unmarshal(body, &getTracksResult)
	if err != nil {
		return nil, err
	}

	var tracks = make([]model.Track, len(getTracksResult.Items))
	for itemIdx, item := range getTracksResult.Items {
		artists := item.Track.Artists.ToString()

		tracks[itemIdx] = model.Track{
			Name:   item.Track.Name,
			Artist: artists,
			Link:   item.Track.ExternalUrls.Spotify,
		}
	}

	return tracks, nil
}
