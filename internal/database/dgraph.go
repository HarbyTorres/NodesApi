package database

import (
	"log"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgo/v200"
)

type Dgraph struct {
	// contains filtered or unexported fields
}

func newClient(client ...api.DgraphClient) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	conn, err := dgo.DialSlashEndpoint("", "")
	if err != nil {
		log.Fatal(err)
	}

	newClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	defer conn.Close()

	return newClient

}

/*
func (d *Dgraph) Alter(ctx context.Context, op *api.Operation) error {

	return err
}
*/
