package languages

import (
	"fmt"
)

type serviceLangs struct {
	repository LangRepository
	repothird LangFromAPI	
}

func NewService(repoLang LangRepository, repoThird LangFromAPI) LangUseCase {
	return &serviceLangs {
		repository: repoLang,
		repothird: repoThird,
	}
}

func (caseLang *serviceLangs) GetLangs() (*[]Language, error) {
	data := caseLang.repothird.GetLangFromAPI()

	for i := 0; i < len(data); i++ {
		result, err := caseLang.repository.StoreLang(&data[i])
		if err != nil {
			fmt.Print("error")
		}
		fmt.Print(result)
	}

	langs, err := caseLang.repository.GetLang()
	if err != nil {
		return nil, err
	} 
	return langs, nil
}