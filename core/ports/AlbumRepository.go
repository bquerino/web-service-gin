package ports

import "mycompany.com/web-service-gin/core/domain"

type AlbumRepository interface{
	Get(id string) (domain.Album, error)
	Save(album domain.Album) error
}
