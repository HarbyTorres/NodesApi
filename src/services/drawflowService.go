package services

import (
	"apinodos/internal/database"
	"apinodos/src/models"
	"apinodos/src/repository"
	"encoding/json"

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
	if err := json.Unmarshal(response, &mapping); err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (d DrawflowSvc) Create(dr models.CreateDrawflow) ([]byte, error) {

	muttation := repository.CreateDrawflowMuttations()
	payload := muttation.SaveDrawflow()
	response, err := database.NewQuery(payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}
