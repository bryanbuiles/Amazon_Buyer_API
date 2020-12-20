package models

// Transaction struct
type Transaction struct {
	ID         string   `json:"id"`
	BuyerID    string   `json:"buyer_id"`
	IP         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIds []string `json:"product_ids"`
}
