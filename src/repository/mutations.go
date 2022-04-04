package repository

import (
	"apinodos/internal/database"
	"context"
	"log"
)

const q = `
		{
			all(func: anyofterms(name, "Alice Bob")) {
				uid
				balance
			}
		}
	`

func getDrawflows() {
	dgClient := database.NewClient()
	txn := dgClient.NewTxn()
	resp, err := txn.Query(context.Background(), q)

	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// After we get the balances, we have to decode them into structs so that
// we can manipulate the data.
