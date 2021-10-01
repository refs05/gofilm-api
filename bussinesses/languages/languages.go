package languages

import "time"

type Language struct {
	Id        int
	Name      string
	CreatedAt time.Time
	DeletedAt time.Time
}
