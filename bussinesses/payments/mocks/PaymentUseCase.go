// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	payments "gofilm/bussinesses/payments"

	mock "github.com/stretchr/testify/mock"
)

// PaymentUseCase is an autogenerated mock type for the PaymentUseCase type
type PaymentUseCase struct {
	mock.Mock
}

// CheckCart provides a mock function with given fields: CartID
func (_m *PaymentUseCase) CheckCart(CartID int) (*payments.Payment, error) {
	ret := _m.Called(CartID)

	var r0 *payments.Payment
	if rf, ok := ret.Get(0).(func(int) *payments.Payment); ok {
		r0 = rf(CartID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*payments.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(CartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatePayment provides a mock function with given fields: cartID, payment
func (_m *PaymentUseCase) ValidatePayment(cartID int, payment *payments.Payment) (*payments.Payment, error) {
	ret := _m.Called(cartID, payment)

	var r0 *payments.Payment
	if rf, ok := ret.Get(0).(func(int, *payments.Payment) *payments.Payment); ok {
		r0 = rf(cartID, payment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*payments.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *payments.Payment) error); ok {
		r1 = rf(cartID, payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
