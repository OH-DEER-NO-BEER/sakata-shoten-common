package client

import "sakata-shoten-common/entity"

type Pay interface {
	Pay(entity.Client, entity.Item)
}
