package card

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID  int
	Amount               int
	Currency             string
	LastFourDigitsOfCard string
	BankReturnCode       string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {

	stripe.Key = c.Secret

	// create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	// params.AddMetadata("key", "value")

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(cardErrorMessage(stripeErr.Code))
		}
		return nil, msg, err
	}
	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your Card Was Declined ."
	case stripe.ErrorCodeExpiredCard:
		msg = "Your Card Is Expired ."
	case stripe.ErrorCodeIncorrectCVC:
		msg = "You Have Entered Incorrect CVC Code ."
	case stripe.ErrorCodeIncorrectZip:
		msg = "You Have Entered Incorrect Zip/Postal Code ."
	case stripe.ErrorCodeAmountTooLarge:
		msg = "You Have Entered Amount Too Large ."
	case stripe.ErrorCodeAmountTooSmall:
		msg = "You Have Entered Amount Too Small ."
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Your Card Has Insufficient Balance ."
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "You Have Entered Incorrect Postal Code ."
	default:
		msg = "Your Card Was Declined ."
	}
	return msg
}
