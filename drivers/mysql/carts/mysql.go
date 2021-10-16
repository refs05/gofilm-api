package carts

import (
	"gofilm/bussinesses/carts"

	"gorm.io/gorm"
)

type mysqlCartsRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) carts.CartRepository {
	return &mysqlCartsRepository{
		DB: db,
	}
}

func (nr *mysqlCartsRepository) GetCartByUser(userID int) (*carts.Cart, error) {
	cart := carts.Cart{}
	err := nr.DB.Where("UserID = ?", userID).First(&cart).Error
	if err != nil {
		return &carts.Cart{}, err
	}
	return &cart, nil
}

func (nr *mysqlCartsRepository) StoreCart(cart *carts.Cart) (*carts.Cart, error) {
	rec := fromDomain(*cart)

	result := nr.DB.Create(rec)

	if result.Error != nil {
		return cart, result.Error
	}

	res := rec.toDomain()
	return &res, nil
}

func (nr *mysqlCartsRepository) UpdateCart(userID int, cart *carts.Cart) (*carts.Cart, error) {
	err := nr.DB.Model(&cart).Where("UserId = ?", userID).Updates(&cart).Error
	if err != nil {
		return &carts.Cart{}, err
	}
	return cart, err
}

func (nr *mysqlCartsRepository) DeleteCart(userID int) error {
	cart := Carts{}
	
	if err := nr.DB.Unscoped().Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (nr *mysqlCartsRepository) GetCartByID(id int) (*carts.Cart, error) {
	cart := carts.Cart{}
	err := nr.DB.Where("id = ?", id).First(&cart).Error
	if err != nil {
		return &carts.Cart{}, err
	}
	return &cart, nil
}

func (nr *mysqlCartsRepository) ChangeStatus(id int) error {
	cart := carts.Cart{}
	err := nr.DB.Model(&cart).Where("Id = ?", id).Update("Is_pay", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (nr *mysqlCartsRepository) GetFilmUser(id int) (*carts.Cart, error) {
	cart := carts.Cart{}
	err := nr.DB.Where("id = ? AND Is_pay = ?", id, true).First(&cart).Error
	if err != nil {
		return &carts.Cart{}, err
	}
	return &cart, nil
}
