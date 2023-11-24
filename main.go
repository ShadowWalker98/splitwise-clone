package main

import (
	"context"
	"splitwise-clone/databasehandler"
	"splitwise-clone/types"
	"time"
)

func main() {

	cake := types.CreateExpense("Sprite", "Soft drink", 10.0,
		"Satyajit", 30, time.November, "Satyajit")

	client := databasehandler.Connect()

	databasehandler.InsertItem(client, cake)

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
