package pay

import (
	"fmt"
	"sakata-shoten-common/entity"
)

type PayPay struct{}

func (p PayPay) Pay(client entity.Client, item entity.Item) {
	fmt.Println("paid by paypay", client, item)
}

func ByPayPay() Pay {
	return new(PayPay)
}
