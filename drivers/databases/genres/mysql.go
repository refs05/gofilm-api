package genres

import (
	"gofilm/bussinesses/genres"

	"gorm.io/gorm"
)

type mysqlGenresRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) genres.GenreRepository {
	return &mysqlGenresRepository{
		DB: db,
	}
}

func (nr *mysqlGenresRepository) StoreGenre(genre *genres.Genre) (*genres.Genre, error) {
	rec := fromDomain(*genre)
	
	result := nr.DB.Create(rec)

	if result.Error != nil {
		return genre, result.Error
	}

	res := rec.toDomain()

	return &res, nil 
} 

func (nr *mysqlGenresRepository) GetGenre() (*[]genres.Genre, error) {
	var genres []genres.Genre

	err := nr.DB.Unscoped().Find(&genres).Error
	if err != nil {
		return nil, err
	}

	return &genres, nil
}