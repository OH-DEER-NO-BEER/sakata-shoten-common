package admin

import "sakata-shoten-common/entity"

func CollectMoney(client entity.Client) int {
	// collect money per month
	return collectMoney(client)
}

func ResisterItem(item entity.Item) {
	resisterItem(item)
}

func Inventory() {
	// tanaorosi
	inventory()
}

func CalculateSales() int {
	//
	return calculateSales()
}
