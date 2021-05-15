package entity

type Client struct {
	ID    int
	Name  string
	Email string
	Debt  int
}

type Clients []Client
