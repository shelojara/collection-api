package adapter

import (
	"github.com/shelojara/collection-api/model"
	"github.com/shelojara/collection-api/port"
	"gorm.io/gorm"
)

var _ port.ListRepository = (*ListRepository)(nil)

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) Create(list *model.List) error {
	return r.db.Create(list).Error
}
