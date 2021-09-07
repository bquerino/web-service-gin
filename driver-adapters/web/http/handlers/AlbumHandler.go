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

func(hdl *HTTPHandler) Get(c *gin.Context){
	album, err := hdl.albumsService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, album)
}

func (hdl *HTTPHandler) Create(context *gin.Context) {
	body := BodyCreate{}
	context.BindJSON(&body)

	album, err := hdl.albumsService.Create(body.Artist, body.Title, body.Price)
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err})
		return
	}
	context.JSON(200, BuildResponseCreate(album))
}