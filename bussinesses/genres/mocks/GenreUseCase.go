// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	genres "gofilm/bussinesses/genres"

	mock "github.com/stretchr/testify/mock"
)

// GenreUseCase is an autogenerated mock type for the GenreUseCase type
type GenreUseCase struct {
	mock.Mock
}

// GetGenreById provides a mock function with given fields: id
func (_m *GenreUseCase) GetGenreById(id int) (*genres.Genre, error) {
	ret := _m.Called(id)

	var r0 *genres.Genre
	if rf, ok := ret.Get(0).(func(int) *genres.Genre); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*genres.Genre)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGenres provides a mock function with given fields:
func (_m *GenreUseCase) GetGenres() (*[]genres.Genre, error) {
	ret := _m.Called()

	var r0 *[]genres.Genre
	if rf, ok := ret.Get(0).(func() *[]genres.Genre); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]genres.Genre)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
