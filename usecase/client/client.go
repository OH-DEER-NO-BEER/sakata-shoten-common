package client

import (
	"sakata-shoten-common/entity"
	"sakata-shoten-common/usecase/client/pay"
)

func Purchase(client entity.Client, item entity.Item, payer pay.Pay) {
	payer.Pay(client, item)
	purchase(client, item)
}
