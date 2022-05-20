package model

type Product struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	Description    string `json:"description"`
	ImageId        string `json:"image_id"`
	Price          string `json:"price"`
	CurrencyID     int    `json:"currency_id"`
	Rating         string `json:"rating"`
	Specifications string `json:"specifications"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
