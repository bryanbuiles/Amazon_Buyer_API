package database

import (
	"github.com/bryanbuiles/tecnical_interview/internal/logs"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// DgraphClient ...
type DgraphClient struct {
	*dgo.Dgraph
}

func newClient() *DgraphClient {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		logs.Error("New Dgraph client fail " + err.Error())
	}

	client := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	return &DgraphClient{client}

}
