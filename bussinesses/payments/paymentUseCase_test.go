package payments_test

import (
	"gofilm/bussinesses/carts"
	"gofilm/bussinesses/collections"
	"gofilm/bussinesses/payments"
	paymentMock "gofilm/bussinesses/payments/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	paymentRepository paymentMock.PaymentRepository
	paymentUseCase payments.PaymentUseCase
	cartUseCase carts.CartUseCase
	collectionUseCase collections.CollectionUseCase
)

func setup() {
	paymentUseCase = payments.NewService(&paymentRepository, cartUseCase, collectionUseCase)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

// func TestCheckCart(t *testing.T) {
// 	t.Run("valid test | test case 1", func(t *testing.T) {
		// domain := payments.Payment{
		// 	Id: 1,
		// 	Verif: true,
		// 	CartID: 23,
		// 	Cart: payments.Cart{},
		// }
// 		domainCart := payments.Cart{
// 			Id: 1,
// 			Total: 100000,
// 			Is_pay: true,
// 			UserID: 3,
// 		}
// 		paymentRepository.On("GetCartByID", mock.AnythingOfType("int")).Return(&domainCart, nil)

// 		result, err := paymentUseCase.CheckCart(1)

// 		assert.Nil(t, err)
// 		assert.Equal(t, payments.Payment{}, result)
// 	})
	
// }

func TestValidatePayment(t *testing.T) {
	t.Run("Valid Test | test case 1", func(t *testing.T) {
		domain := payments.Payment{
			Id: 1,
			Verif: true,
			CartID: 23,
			Cart: payments.Cart{},
		}
		paymentRepository.On("ChangeStatus", mock.AnythingOfType("*payments.Payment")).Return(&domain, nil).Once()

		result, err := paymentUseCase.ValidatePayment(23, &domain)

		assert.Nil(t, err)
		assert.Equal(t, &domain, result)
	})
}