package lists

import (
	"github.com/shelojara/collection-api/port"
	genv1 "github.com/shelojara/collection-api/proto/gen/v1"
	"github.com/shelojara/collection-api/shared/xptr"
)

type UpdateListStatus struct {
	ListRepository port.ListRepository
}

func (it *UpdateListStatus) Execute(req *genv1.UpdateListStatusRequest) (*genv1.UpdateListStatusResponse, error) {
	status, err := it.ListRepository.GetStatus(port.ListStatusQuery{ByID: xptr.Ptr(req.Id)})
	if err != nil {
		return nil, err
	}

	status.Name = req.Name

	if err = it.ListRepository.UpdateStatus(status, req.Fields); err != nil {
		return nil, err
	}

	return &genv1.UpdateListStatusResponse{
		Id: status.ID,
	}, nil
}
