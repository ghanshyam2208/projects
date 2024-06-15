package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()
		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	}

	t.Run("deposit in wallet should work", func(t *testing.T) {
		wallet := NewWallet()
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("initial balance is zero", func(t *testing.T) {
		wallet := NewWallet()
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw amount should be deducted", func(t *testing.T) {
		wallet := NewWallet()
		wallet.Deposit(Bitcoin(20))
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := NewWallet()
		startingBalance := Bitcoin(20)
		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err != ErrInsufficientFunds {
			t.Error("Insufficient error did not match")
		}

	})
}
