package response

import "gofilm/bussinesses/languages"

type Languages struct {
	Kode string `json:"kode"`
	Name string `json:"name"`
}

func ResponseLangs(domain languages.Language) Languages {
	return Languages{
		Kode:   domain.Kode,
		Name: domain.Name,
	}
}

func NewResponseArray(domainLangs []languages.Language) []Languages {
	var resp []Languages

	for _, value := range domainLangs {
		resp = append(resp, ResponseLangs(value))
	}

	return resp
}
