package lists

import (
	"github.com/google/uuid"
	"github.com/shelojara/collection-api/model"
	"github.com/shelojara/collection-api/port"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
)

type CreateList struct {
	ListRepository port.ListRepository
}

func (r *CreateList) Execute(req *genv1.CreateListRequest) (*genv1.CreateListResponse, error) {
	list := &model.List{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	if err := r.ListRepository.Create(list); err != nil {
		return nil, err
	}

	return &genv1.CreateListResponse{
		Id: list.ID,
	}, nil
}
