package models

// Transaction struct
type Transaction struct {
	UID        string           `json:"uid,omitempty"`
	ID         string           `json:"id,omitempty"`
	BuyerID    []UIDTransaction `json:"buyerID,omitempty"`
	IP         string           `json:"ip,omitempty"`
	Device     string           `json:"device,omitempty"`
	ProductIDs []UIDTransaction `json:"productIDs,omitempty"`
	DType      []string         `json:"dgraph.type,omitempty"`
}

// UIDTransaction ...
type UIDTransaction struct {
	UID   string   `json:"uid"`
	DType []string `json:"dgraph.type,omitempty"`
}
