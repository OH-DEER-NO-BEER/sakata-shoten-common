package client

import (
	"fmt"
	"sakata-shoten-common/entity"
)

func purchase(client entity.Client, item entity.Item) {
	// store log to DB
	fmt.Println("purchase", client, item)
}
