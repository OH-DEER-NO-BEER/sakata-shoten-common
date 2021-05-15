package pay

import (
	"fmt"
	"sakata-shoten-common/entity"
	"sakata-shoten-common/usecase/client"
)

type PayPay struct{}

func (p PayPay) Pay(client entity.Client, item entity.Item) {
	fmt.Println("paid by paypay", client, item)
}

func PayByPayPay() client.Pay {
	return new(PayPay)
}
