package genres

import (
	"fmt"
)

type serviceGenres struct {
	repository GenreRepository
	repothird GenreFromAPI
}

func NewService(repoGenre GenreRepository, repoThird GenreFromAPI) GenreUseCase {
	return &serviceGenres {
		repository: repoGenre,
		repothird: repoThird,
	}
}

func (caseGenre *serviceGenres) GetGenres() (*[]Genre, error) {
	
	data := caseGenre.repothird.GetGenreFromAPI()

    for i := 0; i < len(data); i++ {
		result, err := caseGenre.repository.StoreGenre(&data[i])

		if err != nil {
			fmt.Print("error")
		}
		fmt.Print(result)
    }

	genres, err := caseGenre.repository.GetGenre()
	if err != nil {
		return nil, err
	} 
	return genres, nil
}