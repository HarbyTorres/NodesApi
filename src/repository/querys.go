package repository

import (
	"apinodos/internal/database"
	"apinodos/src/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rawNode models.Node
	_ = json.NewDecoder(r.Body).Decode(&rawNode)
	query := getQuery(rawNode.ID)
	dgClient := database.NewClient
	txn := dgClient().NewTxn()
	resp, err := txn.Query(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Json)
}

func getQuery(uid uint64) string {
	return fmt.Sprintf(getFileWithId, uid)
}

const getFileWithId string = `
{
	node(func: uid(%v)) {
	  uid
	  Code
	  expand(_all_) {
		uid
		expand(_all_)
	  }
	}
  }
  `
