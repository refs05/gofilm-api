package collections

type Collection struct {
	Id     int
	UserID int
	Films   []Film `gorm:"foreignKey:Id"`
}

type Film struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	LanguagesKode   string
}

type CollectionUseCase interface {
	StoreCollection(userID int) (*Collection, error)
	GetCollection(userID int) (*Collection, error)
}

type CollectionRepository interface {
	StoreFilm(userID int, collection *Collection) (*Collection, error)
	GetCollection(userID int) (*Collection, error)
}