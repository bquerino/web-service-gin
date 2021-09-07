package ports

import "mycompany.com/web-service-gin/core/domain"

type AlbumService interface{
	Get(id string) (domain.Album, error)
	Create(title string, artist string, price float64) (domain.Album, error)
}
