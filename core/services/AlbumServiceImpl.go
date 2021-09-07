package services

import (
	"errors"
	"mycompany.com/web-service-gin/core/domain"
	"mycompany.com/web-service-gin/core/ports"
	"mycompany.com/web-service-gin/utils"
)

type AlbumService struct{
	albumsRepository ports.AlbumRepository
	uidGen uidgen.UIDGen
}

func New(albumsRepository ports.AlbumRepository, uidGen uidgen.UIDGen) *AlbumService{
	return &AlbumService{
		albumsRepository: albumsRepository,
		uidGen: uidGen,
	}
}

// Get UseCase
func (srv *AlbumService) Get(id string) (domain.Album, error){
	album, err := srv.albumsRepository.Get(id)
	if err != nil{
		return domain.Album{}, errors.New("get album from repository has failed")
	}

	return album, nil
}

// Create UseCase
func (srv *AlbumService) Create(title string, artist string, price float64) (domain.Album, error){

	if price == 0{
		return domain.Album{}, errors.New("Price is invalid. Please insert a valid value.")
	}

	album := domain.NewAlbum(srv.uidGen.New(), title, artist, price)

	if err := srv.albumsRepository.Save(album); err != nil{
		return domain.Album{}, errors.New("Create album into repository has failed.")
	}
	return album, nil
}