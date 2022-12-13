package payment

import (
	"reflect"
	"testing"
)

func TestDDTPay(t *testing.T) {
	tests := []struct {
		method		Paymentmenthod
		loggedin	bool
		expected	string
	}{
		{method: Creditcard, loggedin: false, expected: "Not Logged In"},
		{method: Creditcard, loggedin: true, expected: "Paid with credit card"},
		{method: Debitcard, loggedin: true, expected: "Paid with debit card"},
		{method: Virtualaccount, loggedin: true, expected: "Paid with virtual account"},
		{method: Ewallet, loggedin: true, expected: "Paid with ewallet"},
	}

	for _, tc := range tests {
		result := Pay(tc.method, tc.loggedin)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("expected \"%s\", but got \"%s\"", tc.expected, result)
		}
	}
}
