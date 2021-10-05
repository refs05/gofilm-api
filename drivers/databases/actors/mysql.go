package actors

import 	(
	"gofilm/bussinesses/actors"
	"gorm.io/gorm"
)

type mysqlActorsRepository struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) actors.ActorRepository {
	return &mysqlActorsRepository{
		DB: db,
	}
}

func (nr *mysqlActorsRepository) StoreActor(actor *actors.Actor) (*actors.Actor, error) {
	rec := fromDomain(*actor)
	
	result := nr.DB.Create(rec)

	if result.Error != nil {
		return actor, result.Error
	}

	res := rec.toDomain()

	return &res, nil 
} 

func (nr *mysqlActorsRepository) GetActor() (*[]actors.Actor, error) {
	var actors []actors.Actor

	err := nr.DB.Unscoped().Find(&actors).Error
	if err != nil {
		return nil, err
	}

	return &actors, nil
}