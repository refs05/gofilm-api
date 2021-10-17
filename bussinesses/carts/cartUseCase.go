package carts

import (
	//"fmt"
	"gofilm/bussinesses"
	"gofilm/bussinesses/films"
)

type serviceCarts struct {
	repository CartRepository
	repoFilm films.FilmUseCase

}

func NewService(repoCart CartRepository, repoFilm films.FilmUseCase) CartUseCase {
	return &serviceCarts {
		repository: repoCart,
		repoFilm: repoFilm,
	}
}

func (caseCart *serviceCarts) GetCartByUser(userID int) (*Cart, error) {
	if userID <= 0 {
		return &Cart{}, bussinesses.ErrIDNotFound
	}

	result, err := caseCart.repository.GetCartByUser(userID)
	if err != nil {
		return &Cart{}, err
	}

	return result, nil
}

func (caseCart *serviceCarts) UpdateCartUser(userID int, cart *Cart) (*Cart, error) {
	result, err := caseCart.repository.UpdateCart(userID, cart)
	if err != nil {
		return &Cart{}, err
	}
	return result, nil
}

func (caseCart *serviceCarts) DeleteCartUser(userID int) error {
	return caseCart.repository.DeleteCart(userID)
}

func (caseCart *serviceCarts) CreateCartUser(cart *Convert) (*Cart, error) {
	total := 0

	var film []Film
	for _, value := range cart.Films {
		filmdetail, _ := caseCart.repoFilm.GetFilmByID(value)

		cartFilm := Film{
			Id: filmdetail.Id,
			Title: filmdetail.Title,
			Description: filmdetail.Description,
			ReleaseDate: filmdetail.ReleaseDate,
			Rating: filmdetail.Rating,
			Price: filmdetail.Price,
			Adult: filmdetail.Adult,
			LanguagesKode: filmdetail.LanguagesKode,
		}
		film = append(film, Film(cartFilm))
		total += cartFilm.Price
	}

	detailCart := Cart{
		Id: cart.Id,
		Total: total,
		Is_pay: cart.Is_pay,
		UserID: cart.UserID,
	}
	detailCart.Films = film

	result, err := caseCart.repository.StoreCart(&detailCart)
	if err != nil {
		return &Cart{}, err
	}
	return result, nil
}

func (caseCart *serviceCarts) GetCartByID(id int) (*Cart, error) {
	if id <= 0 {
		return &Cart{}, bussinesses.ErrIDNotFound
	}

	result, err := caseCart.repository.GetCartByUser(id)
	if err != nil {
		return &Cart{}, err
	}

	return result, nil
}

func (caseCart *serviceCarts) ChangeStatus(cartID int) error {
	if cartID <= 0 {
		return bussinesses.ErrIDNotFound
	}

	result := caseCart.repository.ChangeStatus(cartID)
	if result != nil {
		return result
	}
	// if result != nil {
	// 	return 
	// }
	return nil
}

func (caseCart *serviceCarts) GetFilmUser(id int) (*Cart, error) {
	if id <= 0 {
		return &Cart{}, bussinesses.ErrIDNotFound
	}

	result, err := caseCart.repository.GetFilmUser(id)
	if err != nil {
		return &Cart{}, err
	}

	return result, nil
}