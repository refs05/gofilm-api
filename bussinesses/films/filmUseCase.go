package films

import (
	"fmt"
)

type serviceFilms struct {
	repository FilmRepository
	repothird FilmFromAPI
}

func NewService(repoFilm FilmRepository, repoThird FilmFromAPI) FilmUseCase {
	return &serviceFilms {
		repository: repoFilm,
		repothird: repoThird,
	}
}

func (caseFilm *serviceFilms) GetPopularFilms() (*[]Film, error) {
	data := caseFilm.repothird.GetFilmFromAPI()

	for i:=0; i<len(data); i++ {
		result, err := caseFilm.repository.StoreFilm(&data[i])

		if err != nil {
			fmt.Print("error")
		}
		fmt.Print(result)
    }

	genres, err := caseFilm.repository.GetFilm()
	if err != nil {
		return nil, err
	} 
	return genres, nil
}
