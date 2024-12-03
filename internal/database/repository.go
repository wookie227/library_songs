package database

import (
	"song_library/internal/models"

	"gorm.io/gorm"
)

type Songs interface {
	CreateSong(song *models.Song) error
	GetAllSongs(filter map[string]interface{}, offset, limit int) ([]models.Song, error)
	GetSongByID(id int) (*models.Song, error)
	UpdateSong(song *models.Song) error
	DeleteSong(id uint) error
}

type Repository struct {
	Songs
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Songs: NewSongRepository(db),
	}
}
