package carts

import "time"

type Cart struct {
	Id     int
	Total  int
	Is_pay bool
	Films  []Film
	UserID int
}

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
	CartID      int
}

type Genre struct {
	Id   int
	Name string
}

type User struct {
	Id        int
	Address   string
	FirstName string
	LastName  string
	Age       int
	NoHp      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CartUseCase interface {
	GetCartByUser(userID int) (*Cart, error)
	// UpdateCartUser(userID int) (*Cart, error)
	// DeleteCart(userID int) (*Cart, error)
	// CreateCart() (*Cart, error)
}

type CartRepository interface {
	GetCartByUser(userID int) (*Cart, error)
	StoreCart(cart *Cart) (*Cart, error)
	//UpdateCart(userID int) (*Cart, error)
	//DeleteCart(userID int) (*Cart, error)
}
