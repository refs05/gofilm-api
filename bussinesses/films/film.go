package films

//import "time"

type Film struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	Genres      []Genre `gorm:"foreignKey:Id"`
	Languages   string
}

type Genre struct {
	Id   int 
	Name string
}

// type Language struct {
// 	Kode string `gorm:"unique"`
// 	Name string
// }

type FilmUseCase interface {
	GetPopularFilms() (*[]Film, error)
}

type FilmRepository interface {
	GetFilm() (*[]Film, error)
	StoreFilm(film *Film) (*Film, error)
}

type FilmFromAPI interface {
	GetFilmFromAPI() []Film
}
