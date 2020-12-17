package gateway

import (
	"encoding/json"
	"net/http"

	"github.com/bryanbuiles/tecnical_interview/api/V1/models"
	"github.com/bryanbuiles/tecnical_interview/internal/logs"
)

// BuyerGateway al methodos of Buyers user
type BuyerGateway interface {
	ConsumerData() ([]models.Consumer, error)
}

// BuyerService ...
type BuyerService struct {
}

const (
	// URL for amazon api
	URL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/"
)

// ConsumerData ...
func (buyer *BuyerService) ConsumerData() ([]models.Consumer, error) {
	res, err := http.Get(URL + "buyers")
	if err != nil {
		logs.Error("http get fail at COnsumerData " + err.Error())
	}
	defer res.Body.Close()
	var _consumer []models.Consumer

	err = json.NewDecoder(res.Body).Decode(&_consumer)
	if err != nil {
		logs.Error("Decode buyers fails " + err.Error())
	}
	return _consumer, err
}
