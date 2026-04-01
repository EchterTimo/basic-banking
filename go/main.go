package main

import "fmt"

func main() {
	bank := NewBank("Sparkasse Berlin")

	acc1 := bank.CreateAccount("Max")
	acc2 := bank.CreateAccount("Alex")

	fmt.Println(acc1)
	fmt.Println(acc2)
	fmt.Println()

	acc1.Deposit(100)
	acc2.Deposit(-10)

	fmt.Println(acc1)
	fmt.Println(acc2)
	fmt.Println()

	acc1.Transfer(20, &acc2)

	fmt.Println(acc1)
	fmt.Println(acc2)
}
