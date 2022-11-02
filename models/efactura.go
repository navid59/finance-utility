package models

/**
*	Invoice Model
 */

type Invoice struct {
	Number         string   `json:"number"`
	ContractNumber string   `json:"contractNumber"`
	IssueDate      string   `json:"issueDate"`
	DueDate        string   `json:"dueDate"`
	DeliveryDate   string   `json:"deliveryDate"`
	Buyer          Buyer    `json:"buyer"`
	Delivery       Delivery `json:"delivery"`
	Product        Product  `json:"product"`
	Comment        string   `json:"comment"`
	IsTestMod      bool     `json:"isTestMod"`
}

type Buyer struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type Delivery struct {
	Address Address `json:"address"`
}

type Product struct {
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
}

type Address struct {
	Street  string `json:"street"`
	County  string `json:"county"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
