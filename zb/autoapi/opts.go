package autoapi

import (
	"github.com/shopspring/decimal"
	"net/url"
)

type Option[T any] func(T)

type Values struct {
	*Account
	url.Values
	URL         string
	usingTrade  bool
	continueErr bool
}

func WithMarket(market string) Option[*Values] {
	return func(t *Values) {
		t.Set("market", market)
	}
}

func WithAccount(account *Account) Option[*Values] {
	return func(t *Values) {
		t.Account = account
	}
}

func WithTrade() Option[*Values] {
	return func(t *Values) {
		t.usingTrade = true
	}
}

func WithContinueErr() Option[*Values] {
	return func(t *Values) {
		t.continueErr = true
	}
}

func WithCurrency(currency string) Option[*Values] {
	return func(t *Values) {
		t.Set("currency", currency)
	}
}

func WithAcctType(acctType AccountType) Option[*Values] {
	return func(t *Values) {
		t.Set("acctType", string(acctType))
	}
}

func WithEnableExpress() Option[*Values] {
	return func(t *Values) {
		t.Set("enableExpress", "true")
	}
}

func WithEnableRepay() Option[*Values] {
	return func(t *Values) {
		t.Set("enableRepay", "true")
	}
}

func WithAmount(amount decimal.Decimal) Option[*Values] {
	return func(t *Values) {
		t.Set("amount", amount.String())
	}
}

func WithCurrencyMarket(market string) Option[*Values] {
	return func(t *Values) {
		t.Set("currency", market)
	}
}

func WithOrderType(orderType OrderType) Option[*Values] {
	return func(t *Values) {
		t.Set("orderType", string(orderType))
	}
}

func WithCustomerOrderId(orderId string) Option[*Values] {
	return func(t *Values) {
		t.Set("customerOrderId", orderId)
	}
}

func WithTradeType(tradeType TradeType) Option[*Values] {
	return func(t *Values) {
		t.Set("tradeType", string(tradeType))
	}
}

func WithPrice(price decimal.Decimal) Option[*Values] {
	return func(t *Values) {
		t.Set("price", price.String())
	}
}
