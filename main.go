package main

import (
	"fmt"
	"time"
)

type ExpenseItem struct {
	Name         string
	Description  string
	Cost         int32
	BoughtBy     string
	Quantity     int32
	Buyers       []string
	MonthlyCycle time.Month
}

func main() {
	buyers := make([]string, 3)
	buyers[0] = "Satyajit"
	bread := ExpenseItem{
		Name:         "Wonder",
		Description:  "Basic white bread",
		Cost:         3,
		BoughtBy:     "Satyajit",
		Quantity:     1,
		Buyers:       buyers,
		MonthlyCycle: time.November,
	}
	fmt.Println(bread.Name)
}
