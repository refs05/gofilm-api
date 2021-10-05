package request

type Film struct {
	Id          int
	Title       string
	Description string
	ReleaseDate string
	Rating      float32
	Price       int
	Adult       bool
	Genres      []Genre 
	Languages   string
}

type Genre struct {
	Id   int  
	Name string 
}