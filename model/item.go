package model

import genv1 "github.com/shelojara/collection-api/proto/gen/v1"

type ItemKind string

var (
	ItemKindUnspecified ItemKind = ItemKind(genv1.Item_KIND_UNSPECIFIED.String())
	ItemKindGame        ItemKind = ItemKind(genv1.Item_KIND_GAME.String())
)

type Item struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Kind ItemKind

	ExternalSource string `gorm:"index:item_external,unique"`
	ExternalID     string `gorm:"index:item_external,unique"`
	ExternalURL    string

	Title       string
	IGDBRating  float64
	TotalRating float64
	Description string
	CoverURL    string
}
