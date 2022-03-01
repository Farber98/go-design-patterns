package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(CASH)
	if err != nil {
		t.Fatal("payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}
func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DEBIT_CARD)
	if err != nil {
		t.Fatal("payment method of type 'Debit card' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodCreditCard(t *testing.T) {
	payment, err := GetPaymentMethod(CREDIT_CARD)
	if err != nil {
		t.Fatal("payment method of type 'credit card' must exist")
	}

	msg := payment.Pay(32.30)
	if !strings.Contains(msg, "paid using credit card") {
		t.Error("The credit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("a payment method with id 20 must return an error")
	}
	t.Log("LOG:", err)
}
