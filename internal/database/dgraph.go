package database

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgo/v200"
	"github.com/joho/godotenv"
)

type Dgraph struct {
	// contains filtered or unexported fields
}

func newClient(client ...api.DgraphClient) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	apiURL, apiKey := getApiConfig()
	conn, err := dgo.DialSlashEndpoint(apiURL, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	//api.NewDgraphClient(conn)
	newClient := dgo.NewDgraphClient()

	defer conn.Close()

	return newClient

}

func getApiConfig() (string, string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiURL := os.Getenv("DG_API_URL")
	apiKey := os.Getenv("DG_API_KEY")

	return apiURL, apiKey
}

/*
func (d *Dgraph) Alter(ctx context.Context, op *api.Operation) error {

	return err
}
*/
