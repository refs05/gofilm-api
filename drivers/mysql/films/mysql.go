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

	price := nr.DB.Model(&Films{}).Where("price = ?", 0).UpdateColumn("price", 100000)
	if price.Error != nil {
		return film, price.Error
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

// func (nr *mysqlFilmsRepository) Update(id int, user *users.User) (*users.User, error) {

// 	err := nr.DB.Model(&user).Where("id = ?", id).Updates(&user).Error;
// 	if err != nil {
// 		return &users.User{}, err
// 	}
// 	return user, err
// }
