package gateway

import (
	"encoding/json"
	"net/http"
)

// DataHandler ...
func (c *WebServices) DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Date := r.URL.Query().Get("date") // query url parameter
	//res, err := c.data.ConsumerData(Date)
	//res, err := c.data.ProductData(Date)
	res, err := c.data.TransactionData(Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error pasing data"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)

}
