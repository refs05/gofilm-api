package actors

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type serviceActors struct {
	repository ActorRepository
}

func NewService(repoActor ActorRepository) ActorUserCase {
	return &serviceActors {
		repository: repoActor,
	}
}

func (ca *serviceActors) GetActor() (*[]Actor, error) {
	for i:=1; i<150000; i++ {
		response, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/person/%v?api_key=8b350baf5225dd7930c5b98d510101bb", i))

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		//jsonString := string(responseData
	
		var responseObject Actor
		json.Unmarshal(responseData, &responseObject)

		domain := ToDomain(responseObject.Id, responseObject.Name, responseObject.Birthday)

		result, _ := ca.repository.StoreActor(domain)
		fmt.Print(result)
	}
	
	actors, err := ca.repository.GetActor()
	if err != nil {
		return nil, err
	} 
	return actors, nil
}