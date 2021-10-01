package categorys

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type serviceCategorys struct {
	repository CategoryRepository
}

func NewService(repoCategory CategoryRepository) CategoryUseCase {
	return &serviceCategorys {
		repository: repoCategory,
	}
}

func (caseCategory *serviceCategorys) GetByID(id int) (*Category, error) {

	//GetfromAPI
	type Response struct {
		Genres []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"genres"`
	}

	response, err := http.Get("https://api.themoviedb.org/3/genre/movie/list?api_key=8b350baf5225dd7930c5b98d510101bb&language=en-US")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
    json.Unmarshal(responseData, &responseObject)

    fmt.Println(responseObject.Genres)
    fmt.Println(len(responseObject.Genres))

    for i := 0; i < len(responseObject.Genres); i++ {
        fmt.Println(append(responseObject.Genres, ))
    }

	//getGenre
	result, err := caseCategory.repository.GetByID(id)

	if err != nil {
		return &Category{}, err
	}

	return result, nil
}