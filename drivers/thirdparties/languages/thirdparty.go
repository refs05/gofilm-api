package languages

import (
	"encoding/json"
	"fmt"
	"gofilm/bussinesses/languages"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ConsumeLang struct {
	httpClient http.Client
}

func NewConsumeAPI() languages.LangFromAPI {
	return &ConsumeLang {
		httpClient: http.Client{},
	}
}

func (conslang *ConsumeLang) GetLangFromAPI() []languages.Language {
	response, err := http.Get("https://api.themoviedb.org/3/configuration/languages?api_key=8b350baf5225dd7930c5b98d510101bb")
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

	var resp []languages.Language
	for i:=0; i< len(responseObject); i++ {
		resp = append(resp, toDomain(responseObject[i].Kode, responseObject[i].Name))
	}
	
	return resp
}