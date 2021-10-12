package films

type Film struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	Genres      []Genre `gorm:"foreignKey:Id"`
	LanguagesKode   string
}

type Genre struct {
	Id   int 
	Name string
}

type Convert struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	Genres      []int `gorm:"foreignKey:Id"`
	LanguagesKode   string
}

type FilmUseCase interface {
	GetPopularFilms() (*[]Film, error)
	GetFilmByID(id int) (*Film, error)
}

type FilmRepository interface {
	GetFilm() (*[]Film, error)
	GetFilmByID(id int) (*Film, error)
	StoreFilm(film *Film) (*Film, error)
}

type FilmFromAPI interface {
	GetFilmFromAPI() []Convert
}

type FilmGenre interface {
	GetGenreById(id int) (*Genre, error)
}