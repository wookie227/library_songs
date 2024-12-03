package services

import (
	"song_library/internal/database"
	"song_library/internal/models"
)

type SongsService struct {
	repo database.Songs
}

func NewSongsService(repo database.Songs) *SongsService {
	return &SongsService{repo: repo}
}

func (s *SongsService) CreateSong(song *models.Song) error {
	return s.repo.CreateSong(song)
}

func (s *SongsService) GetAllSongs(filter map[string]interface{}, offset, limit int) ([]models.Song, error) {
	return s.repo.GetAllSongs(filter, offset, limit)
}

func (s *SongsService) GetSongByID(id int) (*models.Song, error) {
	return s.repo.GetSongByID(id)
}

func (s *SongsService) UpdateSong(song *models.Song) error {
	return s.repo.UpdateSong(song)
}

func (s *SongsService) DeleteSong(id uint) error {
	return s.repo.DeleteSong(id)
}
