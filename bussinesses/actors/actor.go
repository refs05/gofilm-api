package actors

type Actor struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Birthday string	`json:"birthday"`
}

type ActorUserCase interface {
	GetActor() (*[]Actor, error)
}

type ActorRepository interface {
	StoreActor(actor *Actor) (*Actor, error)
	GetActor() (*[]Actor, error)
}

func ToDomain(id int, name string, birthday string) *Actor {
	return &Actor{
		Id: id,
		Name: name,
		Birthday: birthday,
	}
}
