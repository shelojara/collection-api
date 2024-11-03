package items

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/shelojara/collection-api/model"
	"github.com/shelojara/collection-api/port"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/shared/xerrors"
	"github.com/shelojara/collection-api/shared/xptr"
	"gorm.io/gorm"
)

type ImportGame struct {
	ItemRepository port.ItemRepository
	IGDBService    port.IGDBService
}

func (it *ImportGame) Execute(req *genv1.ImportGameRequest) (*genv1.ImportGameResponse, error) {
	games, err := it.IGDBService.Search(req.Title)
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, fmt.Errorf("game not found: %w", xerrors.ErrNotFound)
	}

	// just assume the first result is the correct one.
	game := games[0]

	item, err := it.ItemRepository.Get(port.ItemQuery{
		ByExternalID:     xptr.Ptr(fmt.Sprintf("%d", game.ID)),
		ByExternalSource: xptr.Ptr("igdb"),
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if item == nil {
		item = &model.Item{
			ID:   uuid.New().String(),
			Kind: model.ItemKindGame,

			ExternalID:     fmt.Sprintf("%d", game.ID),
			ExternalSource: "igdb",
			ExternalURL:    game.Url,

			Title:       game.Name,
			IGDBRating:  game.Rating,
			TotalRating: game.TotalRating,
			Description: game.Summary,
			CoverURL:    fmt.Sprintf("https:%s", game.Cover.Url),
		}

		if err := it.ItemRepository.Create(item); err != nil {
			return nil, err
		}
	} else {
		item.Title = game.Name
		item.IGDBRating = game.Rating
		item.TotalRating = game.TotalRating
		item.Description = game.Summary
		item.CoverURL = fmt.Sprintf("https:%s", game.Cover.Url)

		if err := it.ItemRepository.Update(item, []string{
			"title", "igdb_rating", "total_rating", "description", "cover_url",
		}); err != nil {
			return nil, err
		}
	}

	return &genv1.ImportGameResponse{
		Id: item.ID,
	}, nil
}
