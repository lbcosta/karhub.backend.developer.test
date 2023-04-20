package repository

import (
	"karhub.backend.developer.test/src/api/v1/model"
)

type PlaylistRepository interface {
	GetPlaylist(query, token string) (model.Playlist, error)
	GetTracks(playlistId, token string) (model.Tracks, error)
}
