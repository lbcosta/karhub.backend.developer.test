package model

import "karhub.backend.developer.test/src/api/v1/domain"

type Playlist struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

func (p Playlist) ToDomain() *domain.Playlist {
	return &domain.Playlist{
		ID:     p.ID,
		Name:   p.Name,
		Tracks: Tracks(p.Tracks).ToDomain(),
	}
}

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}

func (t Track) ToDomain() *domain.Track {
	return &domain.Track{
		Name:   t.Name,
		Artist: t.Artist,
		Link:   t.Link,
	}
}

type Tracks []Track

func (t Tracks) ToDomain() []domain.Track {
	tracks := make([]domain.Track, len(t))

	for i, track := range t {
		tracks[i] = *track.ToDomain()
	}

	return tracks
}
