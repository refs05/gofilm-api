package languages

type Language struct {
	Kode string `json:"kode"`
	Name string `json:"name"`
}

type LangUseCase interface {
	GetLangs() (*[]Language, error)
}

type LangRepository interface {
	GetLang() (*[]Language, error)
	StoreLang(lang *Language) (*Language, error)
}

type LangFromAPI interface {
	GetLangFromAPI() []Language
}