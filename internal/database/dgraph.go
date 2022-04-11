package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/joho/godotenv"
)

func NewClient() *dgo.Dgraph {
	apiURL, apiKey := getApiConfig()
	conn, err := dgo.DialSlashEndpoint(apiURL, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	dc := api.NewDgraphClient(conn)
	fmt.Println(dc)
	dg := dgo.NewDgraphClient(dc)
	return dg

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
