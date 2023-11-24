package types

import "time"

type ExpenseItem struct {
	Name         string
	Description  string
	Cost         float32
	BoughtBy     string
	Quantity     int32
	Buyers       []string
	MonthlyCycle time.Month
}

func CreateExpense(name string, description string, cost float32, boughtBy string,
	amount int32, month time.Month, b ...string) ExpenseItem {
	buyers := make([]string, 0)
	for _, buyer := range buyers {
		b = append(buyers, buyer)
	}

	expense := ExpenseItem{
		Name:         name,
		Description:  description,
		Cost:         cost,
		BoughtBy:     boughtBy,
		Quantity:     amount,
		Buyers:       buyers,
		MonthlyCycle: month,
	}
	return expense
}
