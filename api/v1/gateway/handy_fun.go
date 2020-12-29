package gateway

import (
	"context"
	"encoding/json"

	"github.com/bryanbuiles/tecnical_interview/api/v1/models"
	"github.com/bryanbuiles/tecnical_interview/internal/database"
	"github.com/bryanbuiles/tecnical_interview/internal/logs"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

type filterConsumerStruct struct {
	UID string `json:"uid"`
	ID  string `json:"id"`
}

type filterDataResponse struct {
	AllData []filterConsumerStruct `json:"allData"`
}

func filterConsumer(DB *database.DgraphClient, consumerData []models.Consumer) ([]models.Consumer, error) {
	ctx := context.Background()
	q := `{
			allData(func: type(Consumer)) {
				uid
				id
			}
		}`

	txn := DB.NewTxn()
	defer txn.Discard(ctx)
	response, err := txn.Query(ctx, q)
	if err != nil {
		logs.Error("Transaction fails at filterConsumer " + err.Error())
		return nil, err
	}
	var consumerExist *filterDataResponse

	err = json.Unmarshal([]byte(response.Json), &consumerExist)
	if err != nil {
		logs.Error("Unmarshall fails at filterConsumer " + err.Error())
		return nil, err
	}
	var _consumers []models.Consumer
	allconsumerExist := consumerExist.AllData
	for _, values := range consumerData {
		for _, ValuesExists := range allconsumerExist {
			if values.ID == ValuesExists.ID {
				values.UID = ValuesExists.UID
				break
			}
		}
		_consumers = append(_consumers, values)
	}
	return _consumers, nil
}

func filterProduct(DB *database.DgraphClient, productData []models.Product) ([]models.Product, error) {
	ctx := context.Background()
	q := `{
			allData(func: type(Product)) {
				uid
				id
			}
		}`

	txn := DB.NewTxn()
	defer txn.Discard(ctx)
	response, err := txn.Query(ctx, q)
	if err != nil {
		logs.Error("Transaction fails at filterProduct " + err.Error())
		return nil, err
	}
	var productExist *filterDataResponse

	err = json.Unmarshal([]byte(response.Json), &productExist)
	if err != nil {
		logs.Error("Unmarshall fails at filterProduct " + err.Error())
		return nil, err
	}
	var _products []models.Product
	allconsumerExist := productExist.AllData
	for _, values := range productData {
		for _, ValuesExists := range allconsumerExist {
			if values.ID == ValuesExists.ID {
				values.UID = ValuesExists.UID
				break
			}
		}
		_products = append(_products, values)
	}
	return _products, nil
}

// DropData Drop database but keep schema
func DropData(DB *database.DgraphClient) error {
	op := api.Operation{DropOp: api.Operation_DATA}
	err := DB.Alter(context.Background(), &op)
	if err != nil {
		return err
	}
	return nil
}
