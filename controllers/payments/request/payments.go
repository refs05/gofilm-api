package request

import "gofilm/bussinesses/payments"

type PayRequest struct {
	CartID int `json:"cart_id"`
}

func ToDomain(request PayRequest) *payments.Payment {
	return &payments.Payment{
		CartID: request.CartID,
	}
}