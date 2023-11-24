package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"splitwise-clone/databasehandler"
	"splitwise-clone/types"
	"time"
)

func main() {
	httpServerInit()
}

func httpServerInit() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.HandleFunc("/readBody", bodyReader)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}

func bodyReader(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Read Body request")
	fmt.Println(req.URL)
}
func getHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	write, err := w.Write([]byte("Hello World"))
	fmt.Println("Bytes written: ", write)
	if err != nil {
		return
	}
}

func getRoot(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	write, err := writer.Write([]byte("Hello from root"))
	fmt.Println("Bytes written: ", write)
	if err != nil {
		return
	}
}

func addExpense(client *mongo.Client) {
	cake := types.CreateExpense("Sprite", "Soft drink", 10.0,
		"Satyajit", 30, time.November, "Satyajit")

	databasehandler.InsertItem(client, cake)
}

func dataOperations() {
	client := databasehandler.Connect()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
