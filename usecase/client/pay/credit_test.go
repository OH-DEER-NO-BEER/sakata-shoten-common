package pay

import (
	"sakata-shoten-common/entity"
	"testing"
	"time"
)

func TestPayByCredit(t *testing.T) {
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
	payer := ByCredit()
	payer.Pay(client, item)
}
