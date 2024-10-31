package port

import "github.com/shelojara/collection-api/model"

type ListRepository interface {
	Create(list *model.List) error
}
