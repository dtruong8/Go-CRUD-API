package customer

type Customer struct {
	ID         int    `json:"user_id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
}

type CustomerStore interface {
	Customer(id int) (Customer, error)
	Customers() ([]Customer, error)
	CreateCustomer(c *Customer) error
	UpdateCustomer(c *Customer) error
	DeleteCustomer(c *Customer) error
}
