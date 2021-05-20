package pay

import (
	"fmt"
	"sakata-shoten-common/entity"
)

type Credit struct{}

func (p Credit) Pay(client entity.Client, item entity.Item) {
	fmt.Println("paid by credit", client, item)
}

func ByCredit() Pay {
	return new(Credit)
}
