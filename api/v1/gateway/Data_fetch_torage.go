package v1

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
	TransactionData(date string, consumerMap map[string]string, productMap map[string]string) error
	GetAllBuyers() (*api.Response, error)
	GetBuyerInfo(id string) (*api.Response, error)
}

// DataBaseService retrieve database conection
type DataBaseService struct {
	DB *database.DgraphClient
}

type filterDataResponse struct {
	AllData []filterConsumerStruct `json:"allData"`
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

	mapConsumer := make(map[string]string)
	for index, values := range _consumer {
		consumer.ID = values.ID
		consumer.Age = values.Age
		consumer.Name = values.Name
		consumer.DType = []string{"Consumer"}
		mapConsumer[values.ID] = ""
		_consumer[index] = consumer
	}

	// filter data
	consumers, err := filterConsumer(D.DB, _consumer)

	// saving data

	pb, err := json.Marshal(consumers)
	if err != nil {
		logs.Error("Consumer marshall fail at ConsumerData " + err.Error())
		return nil, err
	}
	err = SaveData(D.DB, pb)
	if err != nil {
		logs.Error("Consumer Save Data fail " + err.Error())
		return nil, err
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
	mapProduct := make(map[string]string)
	for {
		line, err := resCsv.Read()
		if err == io.EOF { // end of the file
			break
		}
		if line != nil {
			product.ID = line[0]
			product.Name = line[1]
			product.Price, _ = strconv.Atoi(line[2])
			product.DType = []string{"Product"}
			mapProduct[line[0]] = ""
			_products = append(_products, product)
		}
		if err != nil {
			logs.Error("Fail read lines in csv " + err.Error())
			return nil, err
		}
	}
	//filter data
	products, err := filterProduct(D.DB, _products)

	// saving data

	pb, err := json.Marshal(products)
	if err != nil {
		logs.Error("Product marshall fail at ConsumerData " + err.Error())
		return nil, err
	}
	err = SaveData(D.DB, pb)
	if err != nil {
		logs.Error("Product Save Data fail " + err.Error())
		return nil, err
	}
	return mapProduct, nil
}

// TransactionData ...
func (D *DataBaseService) TransactionData(date string, consumerMap map[string]string, productMap map[string]string) error {
	consumerMap, productMap, err := TransactionUIDS(D.DB, consumerMap, productMap)
	if err != nil {
		logs.Error("TransactionUIDS fail " + err.Error())
		return err
	}

	date = conversorToUnix(date)
	res, err := http.Get(URL + "transactions?date=" + date)
	if err != nil {
		logs.Error("http get fail at TransactionDataData " + err.Error())
		return err
	}
	defer res.Body.Close()
	resbythes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("Fail read lines in Transaction " + err.Error())
		return err
	}

	transactions := strings.Split(string(resbythes), "#")
	var transaction models.Transaction
	var transactionsUIDS models.UIDTransaction
	var transactionsUIDSListProduct []models.UIDTransaction
	var _transactions []models.Transaction
	for index, element := range transactions {
		transactionsUIDSListProduct = nil
		if index == 0 {
			continue
		}
		// \x00 null terminator - carriage returns or line-feeds
		transactionElements := strings.Split(element, "\x00")
		transaction.ID = transactionElements[0]
		transactionsUIDS.UID = consumerMap[transactionElements[1]]
		transactionsUIDS.DType = []string{"Consumer"}
		transaction.BuyerID = []models.UIDTransaction{transactionsUIDS}
		transaction.IP = transactionElements[2]
		transaction.Device = transactionElements[3]
		productsList := strings.Split(transactionElements[4][1:len(transactionElements[4])-1], ",")
		for _, values := range productsList {
			transactionsUIDS.UID = productMap[values]
			transactionsUIDS.DType = []string{"Product"}
			transactionsUIDSListProduct = append(transactionsUIDSListProduct, transactionsUIDS)
		}
		transaction.ProductIDs = transactionsUIDSListProduct
		transaction.DType = []string{"Transaction"}
		_transactions = append(_transactions, transaction)
	}
	// filter data
	result, err := filterTransaction(D.DB, _transactions)
	// saving data
	pb, err := json.Marshal(result)
	if err != nil {
		logs.Error("Transaction marshall fail at TransactionData " + err.Error())
		return err
	}
	err = SaveData(D.DB, pb)
	if err != nil {
		logs.Error("Transaction Save Data fail " + err.Error())
		return err
	}
	return nil
}

//GetAllBuyers Get all buyers by endpoint
func (D *DataBaseService) GetAllBuyers() (*api.Response, error) {
	ctx := context.Background()
	q := `{
		allBuyers(func: type(Consumer)) {
			id
			name
			age
		}
	}`
	txn := D.DB.NewTxn()
	defer txn.Discard(ctx)
	response, err := txn.Query(ctx, q)
	if err != nil {
		logs.Error("Transaction fails at GetAllBuyers " + err.Error())
		return nil, err
	}
	return response, nil

}

// GetBuyerInfo get buyer information
func (D *DataBaseService) GetBuyerInfo(id string) (*api.Response, error) {
	ctx := context.Background()
	variables := map[string]string{"$id": id}
	q := `query BuyerData($id: string) {
			BuyerInfo(func: eq(id, $id)) {
				id
				name
				purchases : ~buyerID {
					transactionID : bid as id
					sameIP as ip
					products : productIDs {
						id
						name
						price
					}
				}
			}
			sameIP(func: eq(ip, val(sameIP))) @filter(NOT eq(id, val(bid))) @normalize {
				ip : ip
				buyerID {
					id : id
			  		name : name
			}
		  }
		  recomendations(func: eq(ip, val(sameIP)), first: 2) @normalize {
			productIDs {
				id : id
				name : name
				price : price
			}
		  }
	}`
	txn := D.DB.NewTxn()
	defer txn.Discard(ctx)
	response, err := txn.QueryWithVars(ctx, q, variables)
	if err != nil {
		logs.Error("Transaction fails at GetBuyerInfo " + err.Error())
		return nil, err
	}
	return response, nil
}
