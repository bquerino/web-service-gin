package repositories

import (
	"encoding/json"
	"errors"
	"mycompany.com/web-service-gin/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Album, error){
	if value, ok := repo.kvs[id]; ok{
		album := domain.Album{}
		err:= json.Unmarshal(value, &album)
		if err != nil{
			return domain.Album{}, errors.New("fail to get value from kvs")
		}
		return album, nil
	}
	return domain.Album{}, errors.New("album not found in kvs")
}

func (repo *memkvs) Save(album domain.Album) error{
	bytes, err := json.Marshal(album)
	if err != nil{
		return errors.New("album fails into json string")
	}
	repo.kvs[album.ID] = bytes

	return nil
}