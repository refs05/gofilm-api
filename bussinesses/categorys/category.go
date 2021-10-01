package categorys

import (


	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Genre string
}

type CategoryUseCase interface {
	GetByID(id int) (*Category, error)
}

type CategoryRepository interface {
	GetByID(id int) (*Category, error)
}