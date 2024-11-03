package adapter

import (
	"github.com/shelojara/collection-api/model"
	"github.com/shelojara/collection-api/port"
	"gorm.io/gorm"
)

var _ port.ItemRepository = (*ItemRepository)(nil)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Get(q port.ItemQuery) (*model.Item, error) {
	db := r.db.Model(&model.Item{})

	if q.ByExternalID != nil {
		db = db.Where("external_id = ?", q.ByExternalID)
	}

	if q.ByExternalSource != nil {
		db = db.Where("external_source = ?", q.ByExternalSource)
	}

	item := &model.Item{}
	if err := db.Take(item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r *ItemRepository) Create(item *model.Item) error {
	return r.db.Create(item).Error
}

func (r *ItemRepository) Update(item *model.Item, fields []string) error {
	return r.db.Model(item).Select(fields).Updates(item).Error
}
