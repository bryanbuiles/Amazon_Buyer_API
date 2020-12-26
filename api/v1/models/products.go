package models

// Product struct
type Product struct {
	UID   string   `json:"uid,omitempty"`
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name"`
	Price int      `json:"price"`
	Dtype []string `json:"dgraph.type,omitempty"`
}
