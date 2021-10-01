package films

import "time"

type Film struct {
	Id          int
	Title       string
	CategoryId  int
	ActorId     int
	Languages   string
	Description string
	ReleaseYear time.Time
	Rating      float32
	Price       int
	Duration    time.Duration
}


