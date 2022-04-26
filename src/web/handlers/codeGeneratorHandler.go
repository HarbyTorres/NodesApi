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
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Content-Type", "application/")
	var draw models.DrawflowMap
	encoder := json.NewDecoder(r.Body)
	encoder.Decode(&draw)
	response := c.codeGenerator.GenerateCode(draw)
	_ = json.NewEncoder(w).Encode(response)
}

func CreateCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}
