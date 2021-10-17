package payments

import (
	"fmt"
	"gofilm/bussinesses/carts"
	"gofilm/bussinesses/collections"
)

type servicePayments struct {
	repository PaymentRepository
	repoCart carts.CartUseCase
	repoCollection collections.CollectionUseCase
}

func NewService(repoPayment PaymentRepository, repoCart carts.CartUseCase, repoCollection collections.CollectionUseCase) PaymentUseCase {
	return &servicePayments {
		repository: repoPayment,
		repoCart: repoCart,
		repoCollection: repoCollection,
	}
}

func (casePayment *servicePayments) ValidatePayment(cartID int, payment *Payment) (*Payment, error) {                                                
	result, err := casePayment.repository.ChangeStatus(payment)
	if err != nil {
		return &Payment{}, err
	}
	statusCart := casePayment.repoCart.ChangeStatus(cartID)
	fmt.Print(statusCart)
	storeFilm, err := casePayment.repoCollection.StoreCollection(cartID)
	fmt.Print(storeFilm)
	if err != nil {
		return &Payment{}, err
	}
	
	return result, nil
}

// func (caseCart *servicePayments) CheckCart(CartID int) (*Payment, error) {
// 	result, err := caseCart.repoCart.GetCartByID(CartID)
// 	// result = &carts.Cart{}
// 	// fmt.Print(*result)
// 	if result != nil {
// 		fmt.Print(*result)
// 	}
// 	if err != nil {
// 		return &Payment{}, err
// 	}
// 	// pay := Payment{
// 	// 	CartID: result.Id,
// 	// }
// 	return &Payment{}, nil
// }

//error untuk store film ke koleksi