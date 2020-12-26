package models

// Consumer struct for buyers
type Consumer struct {
	UID   string   `json:"uid,omitempty"`
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Dtype []string `json:"dgraph.type,omitempty"`
}
