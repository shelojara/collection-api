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

func (c *Lists) GetList(ctx context.Context, req *connect.Request[genv1.GetListRequest]) (*connect.Response[genv1.List], error) {
	var list model.List
	if err := c.DB.Preload("Statuses").Where("id = ?", req.Msg.Id).Take(&list).Error; err != nil {
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
