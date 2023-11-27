package main

import "fmt"

func main() {

	a()

	//cake := types.CreateExpense("Sprite", "Soft drink", 10.0,
	//	"Satyajit", 30, time.November, "Satyajit")
	//
	//client := databasehandler.Connect()
	//
	//databasehandler.InsertItem(client, cake)
	//
	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("value from the panic function and given to recover function is: ", r)
		}
	}()
	b()

	a()

}
func a() {
	// 1. The defer function call is evaluated according to the value of the argument when the defer function
	// is first encountered in the code
	i := 0
	i = 3
	defer fmt.Println(i)
	i = 8
}

func b() {
	// The defer func calls are in last in first out order
	for i := 1; i < 5; i = i + 1 {
		defer fmt.Println("i value: ", i)
	}
	c()

}

func c() (c int) {
	// 3. The named parameters in the return value of a function can be modified using a defer func
	tabs := 4

	tabs++
	tabs++
	tabs++
	defer func() { fmt.Println("panic function") }()
	panic(4)
	// tabs should be 7

	// we will get 7 + 2 = 9 as the return value
	// it reminds me of a closure or something in javascript
	return tabs
}
