// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	languages "gofilm/bussinesses/languages"

	mock "github.com/stretchr/testify/mock"
)

// LangFromAPI is an autogenerated mock type for the LangFromAPI type
type LangFromAPI struct {
	mock.Mock
}

// GetLangFromAPI provides a mock function with given fields:
func (_m *LangFromAPI) GetLangFromAPI() []languages.Language {
	ret := _m.Called()

	var r0 []languages.Language
	if rf, ok := ret.Get(0).(func() []languages.Language); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]languages.Language)
		}
	}

	return r0
}
