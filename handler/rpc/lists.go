package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/shelojara/collection-api/adapter"
	"github.com/shelojara/collection-api/devops"
	"github.com/shelojara/collection-api/interactor/lists"
	"github.com/shelojara/collection-api/model"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/proto/gen/v1/genv1connect"
	"gorm.io/gorm"
)

var _ genv1connect.ListsHandler = (*Lists)(nil)

type Lists struct {
	Config *devops.Config
	DB     *gorm.DB
}

func (h *Lists) GetList(ctx context.Context, req *connect.Request[genv1.GetListRequest]) (*connect.Response[genv1.List], error) {
	var list model.List
	if err := h.DB.Preload("Statuses").Where("id = ?", req.Msg.Id).Take(&list).Error; err != nil {
		return nil, err
	}

	var statuses []*genv1.List_Status
	for _, status := range list.Statuses {
		statuses = append(statuses, &genv1.List_Status{
			Id:   status.ID,
			Name: status.Name,
		})
	}

	return connect.NewResponse(&genv1.List{
		Id:          list.ID,
		Name:        list.Name,
		Description: list.Description,
		Statuses:    statuses,
	}), nil
}

func (h *Lists) ListListItems(ctx context.Context, req *connect.Request[genv1.ListListItemsRequest]) (*connect.Response[genv1.ListListItemsResponse], error) {
	db := h.DB.Model(&model.ListItem{})
	db.Where("list_id = ?", req.Msg.ListId)

	total := int64(0)
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	listItems := make([]*model.ListItem, 0)
	if err := db.
		Preload("Item").Preload("Status").
		Offset(int(req.Msg.Offset)).Limit(int(req.Msg.Limit)).
		Find(&listItems).Error; err != nil {
		return nil, err
	}

	protoListItems := make([]*genv1.List_Item, 0)
	for _, listItem := range listItems {
		protoListItems = append(protoListItems, &genv1.List_Item{
			Id: listItem.ID,
			Item: &genv1.Item{
				Id:    listItem.Item.ID,
				Kind:  genv1.Item_Kind(genv1.Item_Kind_value[string(listItem.Item.Kind)]),
				Title: listItem.Item.Title,
			},
			Status: &genv1.List_Status{
				Id:   listItem.Status.ID,
				Name: listItem.Status.Name,
			},
		})
	}

	return connect.NewResponse(&genv1.ListListItemsResponse{
		Total: int32(total),
		Items: protoListItems,
	}), nil
}

func (c *Lists) CreateList(ctx context.Context, req *connect.Request[genv1.CreateListRequest]) (*connect.Response[genv1.CreateListResponse], error) {
	res, err := (&lists.CreateList{
		ListRepository: adapter.NewListRepository(c.DB),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}

func (c *Lists) CreateListStatus(ctx context.Context, req *connect.Request[genv1.CreateListStatusRequest]) (*connect.Response[genv1.CreateListStatusResponse], error) {
	res, err := (&lists.CreateListStatus{
		ListRepository: adapter.NewListRepository(c.DB),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}

func (c *Lists) UpdateList(ctx context.Context, req *connect.Request[genv1.UpdateListRequest]) (*connect.Response[genv1.UpdateListResponse], error) {
	res, err := (&lists.UpdateList{
		ListRepository: adapter.NewListRepository(c.DB),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}

func (c *Lists) UpdateListStatus(ctx context.Context, req *connect.Request[genv1.UpdateListStatusRequest]) (*connect.Response[genv1.UpdateListStatusResponse], error) {
	res, err := (&lists.UpdateListStatus{
		ListRepository: adapter.NewListRepository(c.DB),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}

func (c *Lists) CreateListItem(ctx context.Context, req *connect.Request[genv1.CreateListItemRequest]) (*connect.Response[genv1.CreateListItemResponse], error) {
	res, err := (&lists.CreateListItem{
		ItemRepository: adapter.NewItemRepository(c.DB),
		ListRepository: adapter.NewListRepository(c.DB),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}
