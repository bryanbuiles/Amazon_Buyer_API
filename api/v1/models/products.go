package models

// Product struct
type Product struct {
	UID   string   `json:"uid,omitempty"`
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price int      `json:"price,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}
