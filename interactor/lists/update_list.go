package lists

import (
	"github.com/shelojara/collection-api/port"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/shared/xptr"
)

type UpdateList struct {
	ListRepository port.ListRepository
}

func (it *UpdateList) Execute(req *genv1.UpdateListRequest) (*genv1.UpdateListResponse, error) {
	list, err := it.ListRepository.Get(port.ListQuery{ByID: xptr.Ptr(req.Id)})
	if err != nil {
		return nil, err
	}

	list.Name = req.Name
	list.Description = req.Description

	if err = it.ListRepository.Update(list, req.Fields); err != nil {
		return nil, err
	}

	return &genv1.UpdateListResponse{
		Id: list.ID,
	}, nil
}
