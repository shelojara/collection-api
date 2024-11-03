package rpc

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"github.com/agnivade/levenshtein"
	"github.com/shelojara/collection-api/adapter"
	"github.com/shelojara/collection-api/devops"
	"github.com/shelojara/collection-api/handler/rpc/mapper"
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

	return connect.NewResponse(mapper.ItemToProto(item)), nil
}

func (h *Items) SearchItems(ctx context.Context, req *connect.Request[genv1.SearchItemsRequest]) (*connect.Response[genv1.SearchItemsResponse], error) {
	db := h.DB.Model(&model.Item{})

	db = db.Where("kind = ?", req.Msg.Kind)

	// search by title using full text search.
	db = db.Where(`to_tsvector(title) @@ to_tsquery(?)`, strings.Join(strings.Split(req.Msg.Query, " "), " & "))

	// search by title using ILIKE.
	db = db.Or("title ILIKE ?", fmt.Sprintf("%%%s%%", req.Msg.Query))

	items := make([]*model.Item, 0)
	if err := db.Limit(20).Find(&items).Error; err != nil {
		return nil, err
	}

	// sort by levenshtein distance.
	slices.SortFunc(items, func(a, b *model.Item) int {
		ad := levenshtein.ComputeDistance(a.Title, req.Msg.Query)
		bd := levenshtein.ComputeDistance(b.Title, req.Msg.Query)

		return bd - ad
	})

	protoItems := make([]*genv1.Item, 0)
	for _, item := range items {
		protoItems = append(protoItems, mapper.ItemToProto(item))
	}

	return connect.NewResponse(&genv1.SearchItemsResponse{
		Items: protoItems,
	}), nil
}

func (h *Items) ImportGame(ctx context.Context, req *connect.Request[genv1.ImportGameRequest]) (*connect.Response[genv1.ImportGameResponse], error) {
	res, err := (&items.ImportGame{
		ItemRepository: adapter.NewItemRepository(h.DB),
		IGDBService:    adapter.NewIGDBService(ctx, h.Config),
	}).Execute(req.Msg)

	return connect.NewResponse(res), err
}
