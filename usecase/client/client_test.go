package client

import (
	"sakata-shoten-common/entity"
	"sakata-shoten-common/usecase/client/pay"
	"testing"
	"time"
)

func TestClientPurchase(t *testing.T) {
	client := entity.Client{
		ID:    0,
		Name:  "test man",
		Email: "a@com",
		Debt:  100}
	item := entity.Item{
		Id:    0,
		Name:  "test item",
		Price: 100,
		Time:  time.Now(),
	}

	Purchase(client, item, pay.ByPayPay())
	Purchase(client, item, pay.ByCredit())
}
