package pay

import (
	"fmt"
	"sakata-shoten-common/entity"
	"sakata-shoten-common/usecase/client"
)

type Credit struct{}

func (p Credit) Pay(client entity.Client, item entity.Item) {
	fmt.Println("paid by paypay", client, item)
}

func PayByCredit() client.Pay {
	return new(Credit)
}
