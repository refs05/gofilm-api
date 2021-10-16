package collections

import (
	"gofilm/bussinesses"
	"gofilm/bussinesses/carts"
)

type serviceCollections struct {
	repository CollectionRepository
	repoCart carts.CartUseCase
}

func NewService(repoCollection CollectionRepository, repoCart carts.CartUseCase) CollectionUseCase {
	return &serviceCollections {
		repository: repoCollection,
		repoCart: repoCart,
	}
}

func (caseCollection *serviceCollections) StoreCollection(cartID int) (*Collection, error) {
	if cartID <= 0 {
		return &Collection{}, bussinesses.ErrIDNotFound
	}

	check, err := caseCollection.repoCart.GetFilmUser(cartID)
	if err != nil {
		return &Collection{}, err
	}

	
	detail := Collection{}
	//detail.Films = check.Films
	var film []Film
	for _, value := range check.Films {
		film = append(film, Film(value))
	} 
	detail.Films = film
	detail.UserID = check.UserID
	result, err := caseCollection.repository.StoreFilm(cartID, &detail)
	if err != nil {
		return &Collection{}, err
	}
	return result, nil
}

func (caseCollection *serviceCollections) GetCollection(userID int) (*Collection, error) {
	if userID <= 0 {
		return &Collection{}, bussinesses.ErrIDNotFound
	}

	result, err := caseCollection.repository.GetCollection(userID)
	if err != nil {
		return &Collection{}, err
	}

	return result, nil
}