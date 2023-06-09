// Code generated by mockery v2.25.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "karhub.backend.developer.test/src/api/v1/model"
)

// PlaylistRepository is an autogenerated mock type for the PlaylistRepository type
type PlaylistRepository struct {
	mock.Mock
}

// GetPlaylist provides a mock function with given fields: query, token
func (_m *PlaylistRepository) GetPlaylist(query string, token string) (model.Playlist, error) {
	ret := _m.Called(query, token)

	var r0 model.Playlist
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (model.Playlist, error)); ok {
		return rf(query, token)
	}
	if rf, ok := ret.Get(0).(func(string, string) model.Playlist); ok {
		r0 = rf(query, token)
	} else {
		r0 = ret.Get(0).(model.Playlist)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(query, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTracks provides a mock function with given fields: playlistId, token
func (_m *PlaylistRepository) GetTracks(playlistId string, token string) (model.Tracks, error) {
	ret := _m.Called(playlistId, token)

	var r0 model.Tracks
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (model.Tracks, error)); ok {
		return rf(playlistId, token)
	}
	if rf, ok := ret.Get(0).(func(string, string) model.Tracks); ok {
		r0 = rf(playlistId, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Tracks)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(playlistId, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPlaylistRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPlaylistRepository creates a new instance of PlaylistRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlaylistRepository(t mockConstructorTestingTNewPlaylistRepository) *PlaylistRepository {
	mock := &PlaylistRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
