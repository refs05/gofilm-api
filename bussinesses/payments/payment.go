package payments

import (

)

type Payment struct {
	Id int `gorm:"primaryKey"`
	Verif bool
	CartID int
	Cart Cart
}

type Cart struct {
	Id     int `gorm:"primaryKey"`
	Total  int
	Is_pay bool
	UserID int
}

type PaymentUseCase interface {
	ValidatePayment(cartID int, payment *Payment) (*Payment, error)
	CheckCart(CartID int) (*Payment, error)
}

type PaymentRepository interface {
	ChangeStatus(payment *Payment) (*Payment, error)
}