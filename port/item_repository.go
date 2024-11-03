package port

import (
	"github.com/shelojara/collection-api/model"
)

type ItemRepository interface {
	Get(q ItemQuery) (*model.Item, error)
	Create(item *model.Item) error

	Update(item *model.Item, fields []string) error
}

type ItemQuery struct {
	ByID             *string
	ByExternalID     *string
	ByExternalSource *string
}
