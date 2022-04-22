package services

import (
	"apinodos/internal/database"
	"apinodos/src/models"
	"apinodos/src/repository"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

type DrawflowService interface {
	create(drawflow models.CreateDrawflow) (*api.Response, error)
}

type DrawflowSvc struct {
}

func (d DrawflowSvc) GetAll() ([]byte, error) {

	muttation := repository.CreateDrawflowMuttations()

	payload := muttation.GetDrawflows()

	response, err := database.NewQuery(payload)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d DrawflowSvc) Create(dr models.CreateDrawflow) ([]byte, error) {

	muttation := repository.CreateDrawflowMuttations()

	payload := muttation.SaveDrawflow()

	response, err := database.NewQuery(payload)

	if err != nil {
		return nil, err
	}

	return response, nil

	/*
		draw, err := json.Marshal(dr)
		if err != nil {
			panic(err)
		}*/
	/*
		dgClient := database.NewClient()
		txn := dgClient.NewTxn()
		mutDraw := &api.Mutation{
			CommitNow: true,
			SetJson:   draw,
		}

		resp, err := txn.Mutate(context.Background(), mutDraw)
	*/
}
