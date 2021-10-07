package films

import (
	"fmt"
	"gofilm/bussinesses/genres"
)

type serviceFilms struct {
	repository FilmRepository
	repothird FilmFromAPI
	repoGenre genres.GenreUseCase
}

func NewService(repoFilm FilmRepository, repoThird FilmFromAPI, repoGenre genres.GenreUseCase) FilmUseCase {
	return &serviceFilms {
		repository: repoFilm,
		repothird: repoThird,
		repoGenre: repoGenre,
	}
}

func (caseFilm *serviceFilms) GetPopularFilms() (*[]Film, error) {
	data := caseFilm.repothird.GetFilmFromAPI()
	

	for i:=0; i<len(data); i++ {
		film := Film{
			Id: data[i].Id,
			Title: data[i].Title,
			Description: data[i].Description,
			ReleaseDate: data[i].ReleaseDate,
			Rating: data[i].Rating,
			Price: data[i].Price,
			Adult: data[i].Adult,
			LanguagesKode: data[i].LanguagesKode,
		}
		
		var genre []Genre
		for _ , value := range data[i].Genres {
			
			gen, _ := caseFilm.repoGenre.GetGenreById(value)

			filmGenre := genres.Genre{
				Id: gen.Id,
				Name: gen.Name,
			}
			genre = append(genre, Genre(filmGenre))
		}
		film.Genres = genre

		result, err := caseFilm.repository.StoreFilm(&film)

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
