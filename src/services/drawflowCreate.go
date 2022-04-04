package services

import (
	"apinodos/src/models"
)

type DrawflowCreateService interface {
	Create(drawflow *models.CreateDrawflow) (*models.Drawflow, error)
}

type DrawflowCreateSvc struct {
	DrawflowCreateService
}

//func NewDrawflowCreateSvc() DrawflowCreateService {}
