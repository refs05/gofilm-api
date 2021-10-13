package response

import "gofilm/bussinesses/payments"

type Payment struct {
	Id int 
	Verif bool
	CartID int
}

func FromDomain(domain payments.Payment) Payment {
	return Payment{
		Id: domain.Id,
		Verif: domain.Verif,
		CartID: domain.CartID,
	}
}

