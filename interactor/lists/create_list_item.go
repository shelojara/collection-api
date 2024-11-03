package lists

import (
	"github.com/google/uuid"
	"github.com/shelojara/collection-api/model"
	"github.com/shelojara/collection-api/port"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/shared/xptr"
)

type CreateListItem struct {
	ListRepository port.ListRepository
	ItemRepository port.ItemRepository
}

func (it *CreateListItem) Execute(req *genv1.CreateListItemRequest) (*genv1.CreateListItemResponse, error) {
	item, err := it.ItemRepository.Get(port.ItemQuery{
		ByID: xptr.Ptr(req.ItemId),
	})
	if err != nil {
		return nil, err
	}

	listItem := &model.ListItem{
		ID:       uuid.New().String(),
		ListID:   req.ListId,
		ItemID:   item.ID,
		StatusID: req.StatusId,
	}

	if err = it.ListRepository.CreateItem(listItem); err != nil {
		return nil, err
	}

	return &genv1.CreateListItemResponse{
		Id: listItem.ID,
	}, nil
}
