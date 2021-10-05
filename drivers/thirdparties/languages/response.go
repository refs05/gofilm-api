package languages

import "gofilm/bussinesses/languages"

type Response []struct {
	Kode string `json:"iso_639_1"`
	Name string `json:"english_name"`
}

func toDomain(kode string, name string) languages.Language {
	return languages.Language{
		Kode: kode,
		Name: name,
	}
}