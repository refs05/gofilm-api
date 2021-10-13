package payments

import "gofilm/bussinesses/payments"

type Payments struct {
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

func (rec *Payments) toDomain() payments.Payment {
	return payments.Payment{
		Id: rec.Id,
		Verif: rec.Verif,
		CartID: rec.CartID,
		Cart: payments.Cart(rec.Cart),
	}
}

func fromDomain(paymentDomain payments.Payment) *Payments {
	return &Payments{
		Id: paymentDomain.Id,
		Verif: paymentDomain.Verif,
		CartID: paymentDomain.CartID,
		Cart: Cart(paymentDomain.Cart),
	}
}