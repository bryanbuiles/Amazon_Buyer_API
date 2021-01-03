package models

// Consumer struct for buyers
type Consumer struct {
	UID   string   `json:"uid,omitempty"`
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   int      `json:"age,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

//ChannelStrutc ...
type ChannelStrutc struct {
	MapHash map[string]string
	DType   string
	Err     error
}
