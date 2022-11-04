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
	Value          string   `json:"value"`
	Tva            string   `json:"tva"`
	Total          string   `json:"total"`
	Buyer          Buyer    `json:"buyer"`
	Delivery       Delivery `json:"delivery"`
	Product        Product  `json:"product"`
	Comment        string   `json:"comment"`
	TestMod        bool     `json:"testMod"`
}

type Buyer struct {
	Name    string  `json:"name"`
	Cif     string  `json:"cif"`
	Cui     string  `json:"cui"`
	Iban    string  `json:"iban"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type Delivery struct {
	Date    string  `json:"date"`
	Address Address `json:"address"`
}

type Product struct {
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Value       string `json:"value"`
	Description string `json:"description"`
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
