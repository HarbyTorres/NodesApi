package services

import (
	"apinodos/internal/database"
	"apinodos/src/models"
	"apinodos/src/repository"
	"encoding/json"
	"strings"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

type DrawflowService interface {
	create(drawflow models.CreateDrawflow) (*api.Response, error)
}

type DrawflowSvc struct {
}

func (d DrawflowSvc) GetAll() (models.DgraphMapping, error) {
	var mapping models.DgraphMapping

	muttation := repository.CreateDrawflowMuttations()
	payload := muttation.GetDrawflows()
	response, err := database.NewQuery(payload)
	if err != nil {
		return mapping, err
	}
	res := []byte(process(string(response)))

	if err := json.Unmarshal(res, &mapping); err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (d DrawflowSvc) Create(dr models.CreateDrawflow) ([]byte, error) {

	muttation := repository.CreateDrawflowMuttations()
	payload := muttation.SaveDrawflow(dr)
	response, err := database.NewQuery(payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func process(str string) string {
	str2 := strings.ReplaceAll((strings.ReplaceAll(string(str), "\"body\":[\"", "\"body\":")), "}}}}}\"]", "}}}}}")
	str2 = strings.ReplaceAll(str2, "'", "\"")

	return str2
}
