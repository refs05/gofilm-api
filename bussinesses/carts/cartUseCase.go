package carts

import "gofilm/bussinesses"

type serviceCarts struct {
	repository CartRepository
}

func NewService(repoCart CartRepository) CartUseCase {
	return &serviceCarts {
		repository: repoCart,
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