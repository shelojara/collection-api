package port

import "github.com/shelojara/collection-api/shared/igdb"

type IGDBService interface {
	Search(title string) ([]*igdb.Game, error)
}
