package payment

type Paymentmenthod int64

const (
	Creditcard Paymentmenthod = iota
	Debitcard
	Virtualaccount
	Ewallet
)

func Pay(method Paymentmenthod, loggedin bool) string {
	if !loggedin {
		return "Not Logged In"
	}
	switch method {
	case Creditcard:
		return "Paid with credit card"
	case Debitcard:
		return "Paid with debit card"
	case Virtualaccount:
		return "Paid with virtual account"
	case Ewallet:
		return "Paid with ewallet"
	}

	return "Payment failed"
}
