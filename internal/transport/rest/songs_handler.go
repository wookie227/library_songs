package rest

import (
	"net/http"
	"song_library/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

func (h *Handler) CreateSong(c *gin.Context) {
	var input models.Song
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.services.CreateSong(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func (h *Handler) GetSongs(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	filter := map[string]interface{}{}
	if group := c.Query("group"); group != "" {
		filter["group"] = group
	}
	if song := c.Query("song"); song != "" {
		filter["song"] = song
	}

	songs, err := h.services.GetAllSongs(filter, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, songs)
}

func (h *Handler) UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Song
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = uint(id)
	if err := h.services.UpdateSong(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, input)
}

func (h *Handler) DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.services.DeleteSong(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
}

func (h *Handler) GetSongById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	song, err := h.services.GetSongByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	var verses []string
	if song.Verses.Status == pgtype.Present {
		for _, v := range song.Verses.Elements {
			verses = append(verses, v.String)
		}
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	totalVerses := len(verses)
	if offset > totalVerses {
		c.JSON(http.StatusOK, gin.H{"verses": []string{}})
		return
	}
	end := offset + limit
	if end > totalVerses {
		end = totalVerses
	}

	paginatedVerses := verses[offset:end]

	response := gin.H{
		"id":     song.ID,
		"group":  song.Group,
		"song":   song.Song,
		"verses": paginatedVerses,
		"total":  totalVerses,
		"offset": offset,
		"limit":  limit,
	}

	c.JSON(http.StatusOK, response)
}
