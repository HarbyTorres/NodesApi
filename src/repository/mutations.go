package repository

import (
	"apinodos/internal/database"
	"context"
	"fmt"
	"log"
)

const q = `
		{
			all(func: anyofterms()) {
				uid
				balance
			}
		}
	`

func SaveDrawflows() {
	dgClient := database.NewClient()
	txn := dgClient.NewTxn()
	resp, err := txn.Query(context.Background(), q)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
