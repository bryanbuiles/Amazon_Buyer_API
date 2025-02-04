package database

import (
	"context"

	"github.com/bryanbuiles/tecnical_interview/internal/logs"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

// DgraphClient ...
type DgraphClient struct {
	*dgo.Dgraph
}

// NewClient start a new connection to Dgraph
func NewClient() *DgraphClient {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		logs.Error("New Dgraph client fail " + err.Error())
	}

	client := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	return &DgraphClient{client}

}

// SetpUpSheme set schema
func SetpUpSheme() {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		logs.Error("Raise Schema client fail " + err.Error())
	}
	defer d.Close()
	client := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	//DropData(client)
	op := &api.Operation{}
	op.Schema = `
		id: string @index(exact) .
		name: string .
		age: int .
		price: int .
		buyerID: [uid] @reverse .
		ip: string @index(hash) .
		device: string @index(hash) .
		productIDs: [uid] @reverse .
		purchases: [uid] .

		type Consumer {
			id
			name
			age
			purchases
		}

		type Product {
			id
			name
			price
		}

		type Transaction {
			id
			buyerID
			ip
			device
			productIDs
		}
	`
	err = client.Alter(context.Background(), op)
	if err != nil {
		logs.Error("Error creacting shema " + err.Error())
	} else {
		logs.Info("scheme succes set")
	}
}

// DropData Drop database but keep schema
func DropData(DB *dgo.Dgraph) error {
	op := api.Operation{DropOp: api.Operation_DATA}
	err := DB.Alter(context.Background(), &op)
	if err != nil {
		return err
	}
	return nil
}
