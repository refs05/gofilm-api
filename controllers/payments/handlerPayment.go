package payments

import (
	"gofilm/bussinesses/payments"
	"gofilm/controllers/payments/request"
	"gofilm/controllers/payments/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	servicePayment payments.PaymentUseCase
}

func NewHandler(paymentServ payments.PaymentUseCase) *Presenter {
	return &Presenter{
		servicePayment: paymentServ,
	}
}

func (handler *Presenter) Validate(e echo.Context) error {
	var req request.PayRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, "Error Noob")
	}

	domain := request.ToDomain(req)
	//_, err := handler.servicePayment.CheckCart(domain.CartID)
	// if err != nil {
	// 	return e.JSON(http.StatusInternalServerError, "Keranjang Anda Tidak Ada")
	// }

	resp, err := handler.servicePayment.ValidatePayment(domain)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Server Interval Error")
	}

	return e.JSON(http.StatusOK, response.FromDomain(*resp))
}