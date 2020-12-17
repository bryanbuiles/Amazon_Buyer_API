package gateway

import (
	"encoding/json"
	"net/http"
)

// ConsumerDateHandler ...
func (c *WebServices) ConsumerDateHandler(w http.ResponseWriter, r *http.Request) {

	res, err := c.buyers.ConsumerData()
	if err != nil {
		m := map[string]interface{}{"msg": "error in create buyer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	_ = json.NewEncoder(w).Encode(res)

}
