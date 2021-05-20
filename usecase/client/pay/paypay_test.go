package pay

import (
	"sakata-shoten-common/entity"
	"testing"
	"time"
)

func TestPayByPayPay(t *testing.T) {
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
	payer := PayByPayPay()
	payer.Pay(client, item)
}
