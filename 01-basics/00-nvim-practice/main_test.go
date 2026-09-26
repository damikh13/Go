package main

import "testing"

func TestDeposit(t *testing.T) {
	acc := &Account{Owner: "Test", Balance: 100}
	acc.Deposit(50)

	if acc.Balance != 150 {
		t.Errorf("expected balance 150, got %d", acc.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name          string
		startBalance  int
		withdrawAmt   int
		expectErr     bool
		expectBalance int
	}{
		{"sufficient funds", 100, 40, false, 60},
		{"exact balance", 100, 100, false, 0},
		{"insufficient funds", 100, 150, true, 100},
		{"zero withdrawal", 100, 0, false, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := &Account{Owner: "Test", Balance: tt.startBalance}
			err := acc.Withdraw(tt.withdrawAmt)

			if tt.expectErr && err == nil {
				t.Errorf("expected an error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if acc.Balance != tt.expectBalance {
				t.Errorf("expected balance %d, got %d", tt.expectBalance, acc.Balance)
			}
		})
	}
}

func TestFibonacciSum(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"zero terms", 0, 0},
		{"one term", 1, 0},
		{"five terms", 5, 7},
		{"ten terms", 10, 88},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FibonacciSum(tt.n)
			if result != tt.expected {
				t.Errorf("FibonacciSum(%d): expected %d, got %d", tt.n, tt.expected, result)
			}
		})
	}
}
