package openapi

import (
	"encoding/json"
	"fmt"
	"gofilm/bussinesses/genres"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ConsumeAPI struct {
	httpClient http.Client 
}

func NewConsumeAPI() genres.GenreFromAPI {
	return &ConsumeAPI {
		httpClient: http.Client{},
	}
}

func (consapi *ConsumeAPI) GetGenreFromAPI() []genres.Genre {
	response, err := http.Get("https://api.themoviedb.org/3/genre/movie/list?api_key=8b350baf5225dd7930c5b98d510101bb&language=en-US")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
    json.Unmarshal(responseData, &responseObject)

	var resp []genres.Genre
	for i:=0; i< len(responseObject.Genres); i++ {
		resp = append(resp, ToDomain(responseObject.Genres[i].Id, responseObject.Genres[i].Name))
	}
	
	return resp
}
