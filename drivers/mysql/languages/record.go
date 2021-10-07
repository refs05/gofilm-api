package languages

import "gofilm/bussinesses/languages"

type Languages struct {
	Kode string `gorm:"unique"`
	Name string
}

func fromDomain(langDomain languages.Language) *Languages {
	return &Languages{
		Kode: langDomain.Kode,
		Name: langDomain.Name,
	}
}

func (rec *Languages) toDomain() languages.Language {
	return languages.Language{
		Kode: rec.Kode,
		Name: rec.Name,
	}
}