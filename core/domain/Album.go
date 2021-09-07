package domain

//album represents data about a record album.
type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

func NewAlbum(id string, title string, artist string, price float64) Album{
	return Album{
		ID: id,
		Title: title,
		Artist: artist,
		Price: price,
	}
}