package pointers_test

import (
	"testing"

	"github.com/11me/go-tests/pointers"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := pointers.Wallet{}

		wallet.Deposit(pointers.Monero(10))
		want := pointers.Monero(10)
		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Monero(20))

		wallet.Withdraw(pointers.Monero(10))
		want := pointers.Monero(10)
		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := pointers.Wallet{}
		startingBalance := pointers.Monero(10)
		wallet.Deposit(startingBalance)

		err := wallet.Withdraw(pointers.Monero(20))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "can not withdraw, insufficient funds")

	})
}

func assertBalance(t testing.TB, wallet pointers.Wallet, want pointers.Monero) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got.String(), want.String())
	}

}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Errorf("wanted an error, but did not get one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
