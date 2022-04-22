package database

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func NewQuery(payload *strings.Reader) ([]byte, error) {
	//apiURL, apiKey := getApiConfig()
	//url := fmt.Sprintf("%v?client=%v", apiURL, apiKey)
	url := "https://blue-surf-590117.us-east-1.aws.cloud.dgraph.io/graphql?postman=MjhmYjlkYmZlMTUxMWY4NGYyYjk4MGMzZWFhODU5Y2Q="

	//payload := strings.NewReader("{\"query\":\"query MyQuery {\\r\\n  getUser(id: \\\"0xfffd8d67db0bee2d\\\") {\\r\\n    id\\r\\n    email\\r\\n    name\\r\\n    password\\r\\n  }\\r\\n}\",\"variables\":{}}")

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

/*
func NewClient() *dgo.Dgraph {
	//apiURL, apiKey := getApiConfig()
		conn, err := dgo.DialCloud("https://blue-surf-590117.us-east-1.aws.cloud.dgraph.io/graphql", "NzVlYWMzMjUxODhmYWI2OGE3MTA3YmMxZGY3ZjNjYmQ=")
		if err != nil {
			log.Fatal(err)
		}
		dc := api.NewDgraphClient(conn)
		dg := dgo.NewDgraphClient(dc)
		fmt.Println(dg)
		return dg

}
*/

func getApiConfig() (string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	apiURL := os.Getenv("DG_API_URL")
	apiKey := os.Getenv("DG_API_KEY")
	return apiURL, apiKey
}
