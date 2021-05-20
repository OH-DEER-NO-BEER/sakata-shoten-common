package client

import (
	"sakata-shoten-common/entity"
	"testing"
	"time"
)

func TestPurchase(t *testing.T) {
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
	purchase(client, item)
}
