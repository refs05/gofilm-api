package payments

import "gofilm/bussinesses/carts"

type servicePayments struct {
	repository PaymentRepository
	repoCart carts.CartUseCase
}

func NewService(repoPayment PaymentRepository, repoCart carts.CartUseCase) PaymentUseCase {
	return &servicePayments {
		repository: repoPayment,
		repoCart: repoCart,
	}
}

func (caseCart *servicePayments) ValidatePayment(payment *Payment) (*Payment, error) {                                                
	result, err := caseCart.repository.ChangeStatus(payment)
	if err != nil {
		return &Payment{}, err
	}
	return result, nil
}

func (caseCart *servicePayments) CheckCart(CartID int) (*Payment, error) {
	_, err := caseCart.repoCart.GetCartByID(CartID)
	if err != nil {
		return &Payment{}, err
	}
	return &Payment{}, nil
}