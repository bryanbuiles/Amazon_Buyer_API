package v1

import (
	"encoding/json"
	"net/http"

	"github.com/bryanbuiles/tecnical_interview/api/v1/models"
	"github.com/go-chi/chi"
)

// DataHandler ...
func (c *WebServices) DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Date := r.URL.Query().Get("date") // query url parameter
	// concurrency
	channel := make(chan models.ChannelStrutc, 2)
	go c.data.ConsumerData(Date, channel)
	go c.data.ProductData(Date, channel)
	data1 := <-channel
	data2 := <-channel
	if data1.Err != nil || data2.Err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"message": "error pasing data at DataHandler"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	productMap := make(map[string]string)
	consumerMap := make(map[string]string)
	if data1.DType == "Products" {
		productMap = data1.MapHash
		consumerMap = data2.MapHash
	} else {
		productMap = data2.MapHash
		consumerMap = data1.MapHash
	}
	err := c.data.TransactionData(Date, consumerMap, productMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"message": "error pasing data in Transaction"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	resMap := map[string]interface{}{"message": "Data saved successfully"}
	_ = json.NewEncoder(w).Encode(resMap)

}

//GetAllBuyerHandler handler to get all buyer
func (c *WebServices) GetAllBuyerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := c.data.GetAllBuyers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"message": "error getting all buyers"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response.Json)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m := map[string]interface{}{"message": "error getting all buyers"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
}

//GetBuyerInfoHandler handler to getbuyer info
func (c *WebServices) GetBuyerInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "buyerID")
	response, err := c.data.GetBuyerInfo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"message": "error getting buyer information"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response.Json)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m := map[string]interface{}{"message": "error getting buyer information"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
}
