package server

import (
	"context"

	types "k8s.io/cri-api/pkg/apis/runtime/v1"

	"github.com/cri-o/cri-o/internal/oci"
)

// ListContainerStats returns stats of all running containers.
func (s *Server) ListContainerStats(ctx context.Context, req *types.ListContainerStatsRequest) (*types.ListContainerStatsResponse, error) {
	ctrList, err := s.ContainerServer.ListContainers(
		func(container *oci.Container) bool {
			return container.StateNoLock().Status != oci.ContainerStateStopped
		},
	)
	if err != nil {
		return nil, err
	}

	if req.GetFilter() != nil {
		cFilter := &types.ContainerFilter{
			Id:            req.GetFilter().GetId(),
			PodSandboxId:  req.GetFilter().GetPodSandboxId(),
			LabelSelector: req.GetFilter().GetLabelSelector(),
		}
		ctrList = s.filterContainerList(ctx, cFilter, ctrList)

		filteredCtrList := []*oci.Container{}

		for _, ctr := range ctrList {
			if filterContainer(ctr.CRIContainer(), cFilter) {
				filteredCtrList = append(filteredCtrList, ctr)
			}
		}

		ctrList = filteredCtrList
	}

	return &types.ListContainerStatsResponse{
		Stats: s.StatsForContainers(ctrList),
	}, nil
}
