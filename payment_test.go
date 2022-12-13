package payment

import (
	"testing"
)

func TestPay_Not_Loggedin(t *testing.T) {
	expected := "Not Logged In"
    result := Pay(Creditcard, false)
    if result != expected {
        t.Errorf("expected \"%s\", but got \"%s\"", expected, result)
    }
}

func TestPay_Creditcard(t *testing.T) {
	expected := "Paid with credit card"
    result := Pay(Creditcard, true)
    if result != expected {
        t.Errorf("expected \"%s\", but got \"%s\"", expected, result)
    }
}

func TestPay_Debitcard(t *testing.T) {
	expected := "Paid with debit card"
    result := Pay(Debitcard, true)
    if result != expected {
        t.Errorf("expected \"%s\", but got \"%s\"", expected, result)
    }
}

func TestPay_Virtualaccount(t *testing.T) {
	expected := "Paid with virtual account"
    result := Pay(Virtualaccount, true)
    if result != expected {
        t.Errorf("expected \"%s\", but got \"%s\"", expected, result)
    }
}

func TestPay_Ewallet(t *testing.T) {
	expected := "Paid with ewallet"
    result := Pay(Ewallet, true)
    if result != expected {
        t.Errorf("expected \"%s\", but got \"%s\"", expected, result)
    }
}
