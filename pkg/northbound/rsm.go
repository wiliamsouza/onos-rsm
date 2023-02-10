// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package northbound

import (
	"context"

	rsmapi "github.com/onosproject/onos-api/go/onos/rsm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"
	"github.com/onosproject/onos-rsm/pkg/nib/rnib"
	"github.com/onosproject/onos-rsm/pkg/nib/uenib"
	"github.com/onosproject/onos-rsm/pkg/slicing"
	"google.golang.org/grpc"
)

func NewService(rnibClient rnib.TopoClient, uenibClient uenib.Client, sm slicing.Manager) service.Service {
	return &Service{
		rnibClient:     rnibClient,
		uenibClient:    uenibClient,
		slicingManager: sm,
	}
}

type Service struct {
	rnibClient     rnib.TopoClient
	uenibClient    uenib.Client
	slicingManager slicing.Manager
}

func (s Service) Register(r *grpc.Server) {
	server := &Server{
		rnibClient:     s.rnibClient,
		uenibClient:    s.uenibClient,
		slicingManager: s.slicingManager,
	}
	rsmapi.RegisterRsmServer(r, server)
}

type Server struct {
	rnibClient     rnib.TopoClient
	uenibClient    uenib.Client
	rsmReqCh       chan *RsmMsg
	slicingManager slicing.Manager
}

func (s Server) CreateSlice(ctx context.Context, request *rsmapi.CreateSliceRequest) (*rsmapi.CreateSliceResponse, error) {
	err := s.slicingManager.HandleNbiCreateSliceRequest(ctx, request, topoapi.ID(request.E2NodeId))
	if err != nil {
		return &rsmapi.CreateSliceResponse{}, err
	}
	return &rsmapi.CreateSliceResponse{
		Ack: &rsmapi.Ack{
			Success: true,
			Cause:   "OK",
		},
	}, nil
}

func (s Server) UpdateSlice(ctx context.Context, request *rsmapi.UpdateSliceRequest) (*rsmapi.UpdateSliceResponse, error) {
	err := s.slicingManager.HandleNbiUpdateSliceRequest(ctx, request, topoapi.ID(request.E2NodeId))
	if err != nil {
		return &rsmapi.UpdateSliceResponse{}, err
	}
	return &rsmapi.UpdateSliceResponse{
		Ack: &rsmapi.Ack{
			Success: true,
			Cause:   "OK",
		},
	}, nil
}

func (s Server) DeleteSlice(ctx context.Context, request *rsmapi.DeleteSliceRequest) (*rsmapi.DeleteSliceResponse, error) {
	err := s.slicingManager.HandleNbiDeleteSliceRequest(ctx, request, topoapi.ID(request.E2NodeId))
	if err != nil {
		return &rsmapi.DeleteSliceResponse{}, err
	}
	return &rsmapi.DeleteSliceResponse{
		Ack: &rsmapi.Ack{
			Success: true,
			Cause:   "OK",
		},
	}, nil
}

func (s Server) SetUeSliceAssociation(ctx context.Context, request *rsmapi.SetUeSliceAssociationRequest) (*rsmapi.SetUeSliceAssociationResponse, error) {
	err := s.slicingManager.HandleNbiSetUeSliceAssociationRequest(ctx, request, topoapi.ID(request.E2NodeId))
	if err != nil {
		return &rsmapi.SetUeSliceAssociationResponse{}, err
	}
	return &rsmapi.SetUeSliceAssociationResponse{
		Ack: &rsmapi.Ack{
			Success: true,
			Cause:   "OK",
		},
	}, nil
}
