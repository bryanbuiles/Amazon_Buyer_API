package gateway

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bryanbuiles/tecnical_interview/api/v1/models"
	"github.com/bryanbuiles/tecnical_interview/internal/database"
	"github.com/bryanbuiles/tecnical_interview/internal/logs"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

// AllDataGateway al methodos of Buyers user
type AllDataGateway interface {
	ConsumerData(date string) (map[string]string, error)
	ProductData(date string) (map[string]string, error)
	TransactionData(date string) ([]models.Transaction, error)
}

// DataBaseService retrieve database conection
type DataBaseService struct {
	DB *database.DgraphClient
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
func (D *DataBaseService) ConsumerData(date string) (map[string]string, error) {

	date = conversorToUnix(date)
	res, err := http.Get(URL + "buyers?date=" + date)

	if err != nil {
		logs.Error("http get fail at COnsumerData " + err.Error())
		return nil, err
	}
	defer res.Body.Close()
	var _consumer []models.Consumer
	var consumer models.Consumer
	err = json.NewDecoder(res.Body).Decode(&_consumer)
	if err != nil {
		logs.Error("Decode buyers fails " + err.Error())
		return nil, err
	}

	for index, values := range _consumer {
		consumer.UID = "_:blank"
		consumer.ID = values.ID
		consumer.Age = values.Age
		consumer.Name = values.Name
		consumer.Dtype = []string{"Consumer"}
		_consumer[index] = consumer
	}

	// filter data
	consumers, err := filterConsumer(D.DB, _consumer)

	// saving data

	mapConsumer := make(map[string]string)
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}
	for _, newvalues := range consumers {
		pb, err := json.Marshal(newvalues)
		if err != nil {
			logs.Error("Consumer marshall fail at ConsumerData " + err.Error())
			return nil, err
		}
		mu.SetJson = pb
		response, err := D.DB.NewTxn().Mutate(ctx, mu)
		if err != nil {
			logs.Error("Error saving Data in Consumer " + err.Error())
			return nil, err
		}
		mapConsumer[newvalues.ID] = response.Uids["blank"]
	}

	return mapConsumer, nil
}

// ProductData ...
func (D *DataBaseService) ProductData(date string) (map[string]string, error) {
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
		if err == io.EOF { // end of the file
			break
		}
		if line != nil {
			product.UID = "_:blank"
			product.ID = line[0]
			product.Name = line[1]
			product.Price, _ = strconv.Atoi(line[2])
			product.Dtype = []string{"Product"}
			_products = append(_products, product)
		}
		if err != nil {
			logs.Error("Fail read lines in csv " + err.Error())
			return nil, err
		}
	}
	mapProduct := make(map[string]string)
	products, err := filterProduct(D.DB, _products)
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}
	for _, newValues := range products {
		pb, err := json.Marshal(newValues)
		if err != nil {
			logs.Error("Product marshall fail at ConsumerData " + err.Error())
			return nil, err
		}
		mu.SetJson = pb
		response, err := D.DB.NewTxn().Mutate(ctx, mu)
		if err != nil {
			logs.Error("Error saving Data in Products " + err.Error())
			return nil, err
		}
		mapProduct[newValues.ID] = response.Uids["blank"]
	}
	return mapProduct, nil
}

// TransactionData ...
func (D *DataBaseService) TransactionData(date string) ([]models.Transaction, error) {
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
		transaction.UID = "_:blank"
		transaction.ID = transactionElements[0]
		transaction.BuyerID = transactionElements[1]
		transaction.IP = transactionElements[2]
		transaction.Device = transactionElements[3]
		transaction.ProductIDs = strings.Split(transactionElements[4][1:len(transactionElements[4])-1], ",")
		transaction.Dtype = []string{"Transaction"}
		_transactions = append(_transactions, transaction)
	}
	return _transactions, nil
}
