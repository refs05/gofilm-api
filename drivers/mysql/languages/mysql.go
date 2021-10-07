package languages

import (
	"gofilm/bussinesses/languages"

	"gorm.io/gorm"
)

type mysqlLanguagesRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) languages.LangRepository {
	return &mysqlLanguagesRepository{
		DB: db,
	}
}

func (nr *mysqlLanguagesRepository) StoreLang(lang *languages.Language) (*languages.Language, error) {
	rec := fromDomain(*lang)

	result := nr.DB.Create(rec)

	if result.Error != nil {
		return lang, result.Error
	}

	res := rec.toDomain()

	return &res, nil
}

func (nr *mysqlLanguagesRepository) GetLang() (*[]languages.Language, error) {
	var langs []languages.Language

	err := nr.DB.Unscoped().Find(&langs).Error
	if err != nil {
		return nil, err
	}

	return &langs, nil
}