package actors

import "gofilm/bussinesses/actors"

type Actors struct {
	Id        int
	Name      string
	Birthday  string
}

func fromDomain(actorDomain actors.Actor) *Actors {
	return &Actors{
		Id:   actorDomain.Id,
		Name: actorDomain.Name,
		Birthday: actorDomain.Birthday,
	}
}

func (rec *Actors) toDomain() actors.Actor {
	return actors.Actor{
		Id:   rec.Id,
		Name: rec.Name,
	}
}