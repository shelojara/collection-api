package lists

import (
	"github.com/google/uuid"
	"github.com/shelojara/collection-api/model"
	"github.com/shelojara/collection-api/port"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
)

type CreateListStatus struct {
	ListRepository port.ListRepository
}

func (it *CreateListStatus) Execute(req *genv1.CreateListStatusRequest) (*genv1.CreateListStatusResponse, error) {
	status := &model.ListStatus{
		ID:     uuid.New().String(),
		ListID: req.ListId,
		Name:   req.Name,
	}

	if err := it.ListRepository.CreateStatus(status); err != nil {
		return nil, err
	}

	return &genv1.CreateListStatusResponse{
		Id: status.ID,
	}, nil
}
