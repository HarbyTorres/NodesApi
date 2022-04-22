package services

import (
	"apinodos/src/models"
	"apinodos/src/repository"
	"strings"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

type DrawflowService interface {
	create(drawflow models.CreateDrawflow) (*api.Response, error)
}

type DrawflowSvc struct {
}

func (d DrawflowSvc) Create(dr models.CreateDrawflow) *strings.Reader {

	muttation := repository.CreateDrawflowMuttations()

	/*
		draw, err := json.Marshal(dr)
		if err != nil {
			panic(err)
		}*/

	payload := muttation.SaveDrawflow()

	/*
		dgClient := database.NewClient()
		txn := dgClient.NewTxn()
		mutDraw := &api.Mutation{
			CommitNow: true,
			SetJson:   draw,
		}

		resp, err := txn.Mutate(context.Background(), mutDraw)
	*/

	return payload
}
