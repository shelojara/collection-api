package mapper

import (
	"github.com/shelojara/collection-api/model"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
)

func ItemToProto(item *model.Item) *genv1.Item {
	return &genv1.Item{
		Id:    item.ID,
		Kind:  genv1.Item_Kind(genv1.Item_Kind_value[string(item.Kind)]),
		Title: item.Title,
	}
}
