package handlers

import (
	"apinodos/src/models"
	"apinodos/src/services"
	"encoding/json"
	"net/http"
)

type CodeGenerator struct {
	codeGenerator services.CodeGenerator
}

func (c *CodeGenerator) GenerateCodeHandler(w http.ResponseWriter, r *http.Request) {
	var draw models.DrawflowMap
	encoder := json.NewDecoder(r.Body)
	encoder.Decode(&draw)
	response := c.codeGenerator.GenerateCode(draw)
	_ = json.NewEncoder(w).Encode(response)
}

func CreateCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}
