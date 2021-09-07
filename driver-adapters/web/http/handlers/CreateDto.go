package handlers

import "mycompany.com/web-service-gin/core/domain"

type BodyCreate struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

type ResponseCreate domain.Album

func BuildResponseCreate(model domain.Album) ResponseCreate {
	return ResponseCreate(model)
}
