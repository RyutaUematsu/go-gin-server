package controller

import (
	"fmt"

	"example.com/youtebe/gin-server/go-youtube-crash-courese/entity"
	"example.com/youtebe/gin-server/go-youtube-crash-courese/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	fmt.Println(video)   // {}
	ctx.BindJSON(&video) // entity.videoの情報を使って取り出す,そしてアドレスに格納される
	fmt.Println(video)   // {title  url}
	c.service.Save(video)
	return video //この戻り値をJSONで返す
}
