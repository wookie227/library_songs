package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/info")
	{
		songs := api.Group("/songs")
		{
			songs.POST("/", h.CreateSong)
			songs.GET("/", h.GetSongs)
			songs.GET("/:id", h.GetSongById)
			songs.PUT("/:id", h.UpdateSong)
			songs.DELETE("/:id", h.DeleteSong)
		}
	}
	return router
}
