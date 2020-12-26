package models

// Transaction struct
type Transaction struct {
	UID        string   `json:"uid,omitempty"`
	ID         string   `json:"id,omitempty"`
	BuyerID    string   `json:"buyerID,omitempty"`
	IP         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIDs []string `json:"productIDs,omitempty"`
	Dtype      []string `json:"dgraph.type,omitempty"`
}
