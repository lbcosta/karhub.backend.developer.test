package service

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"karhub.backend.developer.test/src/api/v1/domain"
	"karhub.backend.developer.test/src/api/v1/model"
	mocks "karhub.backend.developer.test/src/test/mocks/src/api/v1/repository"
	"testing"
)

type PlaylistServiceTestSuite struct {
	suite.Suite
	SomeError          error
	playlistRepository *mocks.PlaylistRepository
	playlistService    PlaylistService
}

func (suite *PlaylistServiceTestSuite) SetupTest() {
	suite.SomeError = errors.New("some error")
	suite.playlistRepository = new(mocks.PlaylistRepository)
	suite.playlistService = NewPlaylistService(suite.playlistRepository)
}

func (suite *PlaylistServiceTestSuite) Test_SearchPlaylists_Success() {
	beerStyles := []string{"IPA"}
	token := "token"

	playlist := model.Playlist{
		Name: "IPA",
	}

	tracks := model.Tracks{
		{
			Name:   "Track 1",
			Artist: "Artist 1",
			Link:   "https://open.spotify.com/track/1",
		},
	}

	playlist.Tracks = tracks

	suite.playlistRepository.On("GetPlaylist", "IPA", token).Return(playlist, nil)
	suite.playlistRepository.On("GetTracks", playlist.ID, token).Return(tracks, nil)

	playlists, err := suite.playlistService.SearchPlaylists(beerStyles, token)

	suite.Nil(err)
	suite.Equal(1, len(playlists))
	suite.Equal(domain.BeerPlaylist{BeerStyle: "IPA", Playlist: *playlist.ToDomain()}, playlists[0])
}

func (suite *PlaylistServiceTestSuite) Test_SearchPlaylists_Error() {
	beerStyles := []string{"IPA"}
	token := "token"

	suite.playlistRepository.On("GetPlaylist", "IPA", token).Return(model.Playlist{}, suite.SomeError)

	playlists, err := suite.playlistService.SearchPlaylists(beerStyles, token)

	suite.Equal(suite.SomeError, err)
	suite.Equal(0, len(playlists))
}

func (suite *PlaylistServiceTestSuite) Test_SearchPlaylists_GetTracksError() {
	beerStyles := []string{"IPA"}
	token := "token"

	playlist := model.Playlist{
		Name: "IPA",
	}

	suite.playlistRepository.On("GetPlaylist", "IPA", token).Return(playlist, nil)
	suite.playlistRepository.On("GetTracks", playlist.ID, token).Return(model.Tracks{}, suite.SomeError)

	playlists, err := suite.playlistService.SearchPlaylists(beerStyles, token)

	suite.Equal(suite.SomeError, err)
	suite.Equal(0, len(playlists))
}

func (suite *PlaylistServiceTestSuite) Test_SearchPlaylists_EmptyBeerStyles() {
	var styles []string

	playlists, err := suite.playlistService.SearchPlaylists(styles, "token")

	suite.Nil(err)
	suite.Equal(0, len(playlists))
}

func TestPlaylistServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PlaylistServiceTestSuite))
}
