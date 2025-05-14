package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("should deposit a balance of 10 BTC", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	t.Run("should withdraw a balance of 5 BTC", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(15)}

		wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	t.Run("should return error because of insufficient funds", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(15)}

		error := wallet.Withdraw(Bitcoin(20))

		want := Bitcoin(15)

		assertBalance(t, wallet, want)

		assertError(t, error, ErrorInsufficientFund)


	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}



func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted error but didn't get one")
	}

	if err != want {
		t.Errorf("want: %s, got: %s", want, err)
	}
}