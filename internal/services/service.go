package services

import (
	"song_library/internal/database"
	"song_library/internal/models"
)

type Songs interface {
	CreateSong(song *models.Song) error
	GetAllSongs(filter map[string]interface{}, offset, limit int) ([]models.Song, error)
	GetSongByID(id int) (*models.Song, error)
	UpdateSong(song *models.Song) error
	DeleteSong(id uint) error
}

type Service struct {
	Songs
}

func NewService(repos *database.Repository) *Service {
	return &Service{
		Songs: NewSongsService(repos.Songs),
	}
}
