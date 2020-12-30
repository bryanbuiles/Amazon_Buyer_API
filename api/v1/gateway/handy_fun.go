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
	allconsumerExist := consumerExist.AllData
	for index, values := range consumerData {
		for _, ValuesExists := range allconsumerExist {
			if values.ID == ValuesExists.ID {
				values.UID = ValuesExists.UID
				consumerData[index] = values
				break
			}
		}
	}
	return consumerData, nil
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
	allconsumerExist := productExist.AllData
	for index, values := range productData {
		for _, ValuesExists := range allconsumerExist {
			if values.ID == ValuesExists.ID {
				values.UID = ValuesExists.UID
				productData[index] = values
				break
			}
		}
	}
	return productData, nil
}

func filterTransaction(DB *database.DgraphClient, transactionData []models.Transaction) ([]models.Transaction, error) {
	ctx := context.Background()
	q := `{
			allData(func: type(Transaction)) {
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
	var transactionExist *filterDataResponse
	err = json.Unmarshal([]byte(response.Json), &transactionExist)
	if err != nil {
		logs.Error("Unmarshall fails at filterTransaction " + err.Error())
		return nil, err
	}

	allTransactionExist := transactionExist.AllData
	for index, values := range transactionData {
		for _, ValuesExists := range allTransactionExist {
			if values.ID == ValuesExists.ID {
				values.UID = ValuesExists.UID
				transactionData[index] = values
				break
			}
		}
	}
	return transactionData, nil
}

// SaveData to save data in dgraph
func SaveData(DB *database.DgraphClient, datajson []byte) error {
	ctx := context.Background()
	txn := DB.NewTxn()
	mu := &api.Mutation{
		SetJson:   datajson,
		CommitNow: true,
	}
	_, err := txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}
	return nil
}

// TransactionUIDS function to take uids for consumer and products
func TransactionUIDS(DB *database.DgraphClient, consumerMap map[string]string, productMap map[string]string) (map[string]string, map[string]string, error) {
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
		logs.Error("Transaction fails at TransactionUIDS " + err.Error())
		return nil, nil, err
	}
	var consumerData *filterDataResponse

	err = json.Unmarshal([]byte(response.Json), &consumerData)
	if err != nil {
		logs.Error("Unmarshall consumer fails at TransactionUIDS " + err.Error())
		return nil, nil, err
	}
	allconsumer := consumerData.AllData
	for keys := range consumerMap {
		for _, values := range allconsumer {
			if keys == values.ID {
				consumerMap[keys] = values.UID
				break
			}
		}
	}
	q = `{
		allData(func: type(Product)) {
			uid
			id
		}
	}`
	var ProductData *filterDataResponse
	response, err = txn.Query(ctx, q)
	if err != nil {
		logs.Error("Transaction fails at TransactionUIDS " + err.Error())
		return nil, nil, err
	}
	err = json.Unmarshal([]byte(response.Json), &ProductData)
	if err != nil {
		logs.Error("Unmarshall product fails at TransactionUIDS " + err.Error())
		return nil, nil, err
	}
	allProducts := ProductData.AllData
	for keys := range productMap {
		for _, values := range allProducts {
			if keys == values.ID {
				productMap[keys] = values.UID
				break
			}
		}
	}
	return consumerMap, productMap, nil
}
