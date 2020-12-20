package gateway

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bryanbuiles/tecnical_interview/api/v1/models"
	"github.com/bryanbuiles/tecnical_interview/internal/logs"
)

// AllDataGateway al methodos of Buyers user
type AllDataGateway interface {
	ConsumerData(date string) ([]models.Consumer, error)
	ProductData(date string) ([]models.Product, error)
	TransactionData(date string) ([]models.Transaction, error)
}

// DataBaseService retrieve database conection
type DataBaseService struct {
}

const (
	// URL for amazon api
	URL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/"
)

func conversorToUnix(date string) string {
	if date == "" {
		return ""
	}
	timeDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		logs.Error("Date fail to parser " + err.Error())
	}
	timeUnix := timeDate.Unix()
	ToString := strconv.FormatInt(timeUnix, 10)
	return ToString
}

// ConsumerData ...
func (DB *DataBaseService) ConsumerData(date string) ([]models.Consumer, error) {

	date = conversorToUnix(date)
	res, err := http.Get(URL + "buyers?date=" + date)

	if err != nil {
		logs.Error("http get fail at COnsumerData " + err.Error())
		return nil, err
	}
	defer res.Body.Close()
	var _consumer []models.Consumer

	err = json.NewDecoder(res.Body).Decode(&_consumer)
	if err != nil {
		logs.Error("Decode buyers fails " + err.Error())
		return nil, err
	}
	return _consumer, nil
}

// ProductData ...
func (DB *DataBaseService) ProductData(date string) ([]models.Product, error) {
	date = conversorToUnix(date)
	res, err := http.Get(URL + "products?date=" + date)

	if err != nil {
		logs.Error("http get fail at productData " + err.Error())
		return nil, err
	}
	defer res.Body.Close()
	resCsv := csv.NewReader(res.Body)
	resCsv.Comma = '\''
	var product models.Product
	var _products []models.Product
	for {
		line, err := resCsv.Read()
		if err == io.EOF {
			break
		}
		if line != nil {
			product.ID = line[0]
			product.Name = line[1]
			product.Price, _ = strconv.Atoi(line[2])
			_products = append(_products, product)
		}
		if err != nil {
			logs.Error("Fail read lines in csv " + err.Error())
			return nil, err
		}
	}
	return _products, nil
}

// TransactionData ...
func (DB *DataBaseService) TransactionData(date string) ([]models.Transaction, error) {
	date = conversorToUnix(date)
	res, err := http.Get(URL + "transactions?date=" + date)

	if err != nil {
		logs.Error("http get fail at TransactionDataData " + err.Error())
		return nil, err
	}
	defer res.Body.Close()
	resbythes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("Fail read lines in Transaction " + err.Error())
		return nil, err
	}
	transactions := strings.Split(string(resbythes), "#")
	var transaction models.Transaction
	var _transactions []models.Transaction
	for index, element := range transactions {

		if index == 0 {
			continue
		}
		// \x00 null terminator - carriage returns or line-feeds
		transactionElements := strings.Split(element, "\x00")

		transaction.ID = transactionElements[0]
		transaction.BuyerID = transactionElements[1]
		transaction.IP = transactionElements[2]
		transaction.Device = transactionElements[3]
		transaction.ProductIds = strings.Split(transactionElements[4][1:len(transactionElements[4])-1], ",")
		_transactions = append(_transactions, transaction)
	}
	return _transactions, nil
}
