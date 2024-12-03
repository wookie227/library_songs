package database

import (
	"song_library/internal/models"

	"gorm.io/gorm"
)

type SongsRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongsRepository {
	return &SongsRepository{db: db}
}

func (r *SongsRepository) CreateSong(song *models.Song) error {
	return r.db.Create(song).Error
}

func (r *SongsRepository) GetAllSongs(filter map[string]interface{}, offset, limit int) ([]models.Song, error) {
	var songs []models.Song
	query := r.db.Model(&models.Song{}).Where(filter).Offset(offset).Limit(limit)
	err := query.Find(&songs).Error
	return songs, err
}

func (r *SongsRepository) GetSongByID(id int) (*models.Song, error) {
	var song models.Song
	err := r.db.First(&song, id).Error
	return &song, err
}

func (r *SongsRepository) UpdateSong(song *models.Song) error {
	return r.db.Save(song).Error
}

func (r *SongsRepository) DeleteSong(id uint) error {
	return r.db.Delete(&models.Song{}, id).Error
}
