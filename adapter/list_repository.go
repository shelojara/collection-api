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

func (r *ListRepository) Get(q port.ListQuery) (*model.List, error) {
	db := r.db.Model(&model.List{})

	if q.ByID != nil {
		db.Where("id = ?", q.ByID)
	}

	list := &model.List{}
	if err := db.Take(list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (r *ListRepository) GetStatus(q port.ListStatusQuery) (*model.ListStatus, error) {
	db := r.db.Model(&model.ListStatus{})

	if q.ByID != nil {
		db.Where("id = ?", q.ByID)
	}

	status := &model.ListStatus{}
	if err := db.Take(status).Error; err != nil {
		return nil, err
	}

	return status, nil
}

func (r *ListRepository) Create(list *model.List) error {
	return r.db.Create(list).Error
}

func (r *ListRepository) CreateStatus(status *model.ListStatus) error {
	return r.db.Create(status).Error
}

func (r *ListRepository) CreateItem(item *model.ListItem) error {
	return r.db.Create(item).Error
}

func (r *ListRepository) Update(list *model.List, fields []string) error {
	return r.db.Model(list).Select(fields).Updates(list).Error
}

func (r *ListRepository) UpdateStatus(status *model.ListStatus, fields []string) error {
	return r.db.Model(status).Select(fields).Updates(status).Error
}
