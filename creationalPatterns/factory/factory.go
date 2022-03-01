package factory

import (
	"errors"
	"fmt"
)

/*
FACTORY:
- Delegate the creation of new instances of structures to a differente part of the program.
- Working at the interface level instead of with concrete implementation.
- Group families of objects to obtain a family object creator.

ANALOGY: Holidays with travel agency.
- You don't deal with hotels and traveling, you just tell the agency the destination.
- Travel agency provides everything you need.
- Travel agency represent a Factory of trips.

EXAMPLE: Factory of payment methods for a shop.
- Implement a payments method factory, to provide us with different ways of paying at a shop.
- We need to have a common method for every payment method, called Pay.
- We need to be able to delegate the creation of paymetns methods of the Factory.
- We need to be able to add more payment methods to the library by just adding it to the factory mthod.

*/

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	CASH        = 1
	DEBIT_CARD  = 2
	CREDIT_CARD = 3
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case CASH:
		return new(CashPM), nil
	case DEBIT_CARD:
		return new(DebitCardPM), nil
	case CREDIT_CARD:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("payment method not recognized"))
	}
}

type CashPM struct {
}

type DebitCardPM struct {
}
type CreditCardPM struct {
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)

}
func (cc *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit card\n", amount)

}
