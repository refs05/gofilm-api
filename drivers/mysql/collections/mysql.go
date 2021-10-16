package collections

import (
	"gofilm/bussinesses/collections"

	"gorm.io/gorm"
)

type mysqlCollectionsRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) collections.CollectionRepository {
	return &mysqlCollectionsRepository{
		DB: db,
	}
}

func (nr *mysqlCollectionsRepository) StoreFilm(userID int, collection *collections.Collection) (*collections.Collection, error) {
	rec := fromDomain(*collection)

	result := nr.DB.Create(rec)
	if result.Error != nil {
		return collection, result.Error
	}

	res := rec.toDomain()
	return &res, nil
}

func (nr *mysqlCollectionsRepository) GetCollection(userID int) (*collections.Collection, error) {
	collection := collections.Collection{}
	err := nr.DB.Where("user_id = ?", userID).First(&collection).Error
	if err != nil {
		return &collections.Collection{}, err
	}
	return &collection, nil
}