package port

import "github.com/shelojara/collection-api/model"

type ListRepository interface {
	Get(q ListQuery) (*model.List, error)
	GetStatus(q ListStatusQuery) (*model.ListStatus, error)

	Create(list *model.List) error
	CreateStatus(status *model.ListStatus) error
	CreateItem(item *model.ListItem) error

	Update(list *model.List, fields []string) error
	UpdateStatus(status *model.ListStatus, fields []string) error
}

type ListQuery struct {
	ByID *string
}

type ListStatusQuery struct {
	ByID *string
}
