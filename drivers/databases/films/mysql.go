package films

import (
	"gofilm/bussinesses/films"

	"gorm.io/gorm"
)

type mysqlFilmsRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) films.FilmRepository {
	return &mysqlFilmsRepository{
		DB: db,
	}
}

func (nr *mysqlFilmsRepository) StoreFilm(film *films.Film) (*films.Film, error) {
	rec := fromDomain(*film)

	result := nr.DB.Create(rec)

	if result.Error != nil {
		return film, result.Error
	}

	res := rec.toDomain()

	return &res, nil
}

func (nr *mysqlFilmsRepository) GetFilm() (*[]films.Film, error) {
	var films []films.Film

	err := nr.DB.Unscoped().Find(&films).Error
	if err != nil {
		return nil, err
	}

	return &films, nil
}