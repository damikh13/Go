package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

// Deposit adds amount to the account balance.
func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

// Withdraw removes amount from the balance if funds are sufficient.
// Returns an error if the withdrawal would overdraw the account.
func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds: balance %d, requested %d", a.Balance, amount)
	}
	a.Balance -= amount
	return nil
}

func FibonacciSum(n int) int {
	sum := 0
	a, b := 0, 1
	for i := 0; i < n; i++ {
		sum += a
		a, b = b, a+b
	}
	return sum
}

func main() {
	acc := &Account{Owner: "Dan", Balance: 100}

	acc.Deposit(50)
	fmt.Println("Balance after deposit:", acc.Balance)

	if err := acc.Withdraw(30); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Balance after withdrawal:", acc.Balance)

	result := FibonacciSum(10)
	fmt.Println("Sum of first 10 Fibonacci numbers:", result)

	if err := acc.Withdraw(9999); err != nil {
		fmt.Println("Error:", err)
	}
}
