package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	balanceAssertion := func(t *testing.T, actual Bitcoin, expected Bitcoin) {
		t.Helper()
		if actual != expected {
			t.Errorf("got '%s' - wanted '%s' ", actual, expected)
		}
	}

	errorAssertion := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Fatal("didn't get an error but wanted one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		expected := Bitcoin(10)
		actual := wallet.Balance()
		balanceAssertion(t, actual, expected)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err, actual := wallet.Withdraw(100)
		balanceAssertion(t, actual, 20)
		errorAssertion(t, err)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		_, actual := wallet.Withdraw(10)
		balanceAssertion(t, actual, 10)
	})

}
