package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/shelojara/collection-api/adapter"
	"github.com/shelojara/collection-api/interactor/lists"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/proto/gen/v1/genv1connect"
	"gorm.io/gorm"
)

var _ genv1connect.ListsHandler = (*Lists)(nil)

type Lists struct {
	DB *gorm.DB
}

func (c *Lists) CreateList(ctx context.Context, req *connect.Request[genv1.CreateListRequest]) (*connect.Response[genv1.List], error) {
	res, err := (&lists.CreateList{
		ListRepository: adapter.NewListRepository(c.DB),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}
