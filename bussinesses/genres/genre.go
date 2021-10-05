package genres

type Genre struct {
	Id int `json:"id"` 
	Name string `json:"name"`
}

type GenreUseCase interface {
	GetGenres() (*[]Genre, error)
}

type GenreRepository interface {
	GetGenre() (*[]Genre, error)
	StoreGenre(genre *Genre) (*Genre, error)
}

type GenreFromAPI interface {
	GetGenreFromAPI() []Genre
}
