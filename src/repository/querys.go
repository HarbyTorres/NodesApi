package repository

import (
	"fmt"
	"net/http"
)

type GetDrawflow struct {
}

func CreateGetDrawflow() *GetDrawflow {
	return &GetDrawflow{}
}

func (p *GetDrawflow) GetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
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
	*/
}

func getQuery(uid int) string {
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
