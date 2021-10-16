package payments

import (
	//"gofilm/bussinesses"
	"gofilm/bussinesses/payments"

	"gorm.io/gorm"
)

type mysqlPaymentsRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) payments.PaymentRepository {
	return &mysqlPaymentsRepository{
		DB: db,
	}
}

func (nr *mysqlPaymentsRepository) ChangeStatus(payment *payments.Payment) (*payments.Payment, error) {
	rec := fromDomain(*payment)
	//if rec.Verif != true {
		rec.Verif = true
	// } else {
	// 	return payment, bussinesses.ErrDuplicateData
	// }
	
	result := nr.DB.Create(rec)
	if result.Error != nil {
		return payment, result.Error
	}
	res := rec.toDomain()

	return &res, nil
	// err := nr.DB.Model(&payment).Where("CartID = ?", CartID).Update("Verif", true).Error
	// if err != nil {
	// 	return &payments.Payment{}, err
	// }
	// return payment, err

}
