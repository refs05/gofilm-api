package request

import "gofilm/bussinesses/collections"

type Collection struct {
	UserID     int `json:"user_id"`
}

// type Film struct {
// 	Id          int
// 	Title       string
// 	Description string
// 	ReleaseDate string
// 	Rating      float32
// 	Price       int
// 	Adult       bool
// 	LanguagesKode   string
// }

func ToDomain(request Collection) *collections.Collection {
	return &collections.Collection{
		UserID: request.UserID,
		// Id: request.Id,
		//Films: request.Films,
	}
}