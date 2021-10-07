package films

import (
	"gofilm/bussinesses/films"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"encoding/json"
)

type ConsumeAPI struct {
	httpClient http.Client
}

func NewConsumeAPI() films.FilmFromAPI {
	return &ConsumeAPI{
		httpClient: http.Client{},
	}
}

func (consapi *ConsumeAPI) GetFilmFromAPI() []films.Convert {
	response, err := http.Get("https://api.themoviedb.org/3/movie/top_rated?api_key=8b350baf5225dd7930c5b98d510101bb&page=1")
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
	
	var resp []films.Convert
	for i:=0; i< len(responseObject.Films); i++ {
		resp = append(resp, ToDomain(responseObject.Films[i]))
	}
	return resp
}