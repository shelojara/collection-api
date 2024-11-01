package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/shelojara/collection-api/adapter"
	"github.com/shelojara/collection-api/devops"
	"github.com/shelojara/collection-api/interactor/items"
	"github.com/shelojara/collection-api/model"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/proto/gen/v1/genv1connect"
	"gorm.io/gorm"
)

var _ genv1connect.ItemsHandler = (*Items)(nil)

type Items struct {
	Config *devops.Config
	DB     *gorm.DB
}

func (c *Items) GetItem(ctx context.Context, req *connect.Request[genv1.GetItemRequest]) (*connect.Response[genv1.Item], error) {
	item := &model.Item{}
	if err := c.DB.Where("id = ?", req.Msg.Id).Take(item).Error; err != nil {
		return nil, err
	}

	return connect.NewResponse(&genv1.Item{
		Id:    item.ID,
		Kind:  genv1.Item_Kind(genv1.Item_Kind_value[string(item.Kind)]),
		Title: item.Title,
	}), nil
}

func (h *Items) ImportGame(ctx context.Context, req *connect.Request[genv1.ImportGameRequest]) (*connect.Response[genv1.ImportGameResponse], error) {
	res, err := (&items.ImportGame{
		ItemRepository: adapter.NewItemRepository(h.DB),
		IGDBService:    adapter.NewIGDBService(ctx, h.Config),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}
