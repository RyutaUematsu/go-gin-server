package service

import "example.com/youtebe/gin-server/go-youtube-crash-courese/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	// 引数の videoを追加
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	//すべてのビデオを返す
	return service.videos
}
