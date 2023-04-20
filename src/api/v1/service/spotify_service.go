package service

import (
	"karhub.backend.developer.test/src/api/v1/domain"
	"karhub.backend.developer.test/src/api/v1/repository"
)

type SpotifyService struct {
	playlistRepository repository.PlaylistRepository
}

func NewSpotifyService(playlistRepository repository.PlaylistRepository) SpotifyService {
	return SpotifyService{playlistRepository: playlistRepository}
}

func (s SpotifyService) SearchPlaylists(beerStyles []string, token string) ([]domain.BeerPlaylist, error) {
	playlists := make([]domain.BeerPlaylist, len(beerStyles))

	for beerStyleIdx, beerStyle := range beerStyles {
		playlist, err := s.playlistRepository.GetPlaylist(beerStyle, token)
		if err != nil {
			return nil, err
		}

		tracks, err := s.playlistRepository.GetTracks(playlist.ID, token)
		if err != nil {
			return nil, err
		}

		playlists[beerStyleIdx] = domain.BeerPlaylist{
			BeerStyle: beerStyle,
			Playlist: domain.Playlist{
				Name:   playlist.Name,
				Tracks: tracks.ToDomain(),
			},
		}
	}

	return playlists, nil
}
