package gateway

import (
	"encoding/json"
	"net/http"
)

// DataHandler ...
func (c *WebServices) DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Date := r.URL.Query().Get("date") // query url parameter
	consumerMap, err := c.data.ConsumerData(Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"message": "error pasing data in buyer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	productMap, err := c.data.ProductData(Date)
	err = c.data.TransactionData(Date, consumerMap, productMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"message": "error pasing data"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	resMap := map[string]interface{}{"msg": "Data saved successfully"}
	_ = json.NewEncoder(w).Encode(resMap)

}
