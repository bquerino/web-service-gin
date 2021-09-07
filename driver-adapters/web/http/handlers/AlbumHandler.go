package handlers

import (
	"github.com/gin-gonic/gin"
	"mycompany.com/web-service-gin/core/ports"
)

//Driver adapter
type HTTPHandler struct {
	albumsService ports.AlbumService
}

func NewHTTPHandler(albumsService ports.AlbumService) *HTTPHandler {
	return &HTTPHandler{
		albumsService: albumsService,
	}
}

func (handler *HTTPHandler) Get(context *gin.Context) {
	album, err := handler.albumsService.Get(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	context.JSON(200, album)
}

func (handler *HTTPHandler) Create(context *gin.Context) {

	// Maps to Dto
	body := BodyCreate{}

	//Bind the context to a JSON structure
	context.BindJSON(&body)

	album, err := handler.albumsService.Create(body.Artist, body.Title, body.Price)

	// Check if you don't have any errors.
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	context.JSON(200, BuildResponseCreate(album))
}
