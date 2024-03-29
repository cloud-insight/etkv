// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3rpc

import (
	"context"
	"fmt"
	"github.com/ucloud/etkv/etkvserver"

	"go.etcd.io/etcd/etcdserver"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/lease"

	"go.uber.org/zap"
)

type LeaseServer struct {
	lg  *zap.Logger
	hdr header
	le  etcdserver.Lessor
}

func NewLeaseServer(s *etkvserver.EtkvServer) pb.LeaseServer {
	return &LeaseServer{lg: s.Cfg.Logger, le: s, hdr: newHeader(s)}
}

func (ls *LeaseServer) LeaseGrant(ctx context.Context, cr *pb.LeaseGrantRequest) (*pb.LeaseGrantResponse, error) {
	resp, err := ls.le.LeaseGrant(ctx, cr)

	if err != nil {
		return nil, togRPCError(err)
	}
	ls.hdr.fill(resp.Header)
	return resp, nil
}

func (ls *LeaseServer) LeaseRevoke(ctx context.Context, rr *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	resp, err := ls.le.LeaseRevoke(ctx, rr)
	if err != nil {
		return nil, togRPCError(err)
	}
	ls.hdr.fill(resp.Header)
	return resp, nil
}

func (ls *LeaseServer) LeaseTimeToLive(ctx context.Context, rr *pb.LeaseTimeToLiveRequest) (*pb.LeaseTimeToLiveResponse, error) {
	resp, err := ls.le.LeaseTimeToLive(ctx, rr)
	defer func() { ls.hdr.fill(resp.Header) }()
	if err == nil {
		return resp, nil
	}

	if err != lease.ErrLeaseNotFound {
		return nil, togRPCError(err)
	} else {
		resp = &pb.LeaseTimeToLiveResponse{
			Header: &pb.ResponseHeader{},
			ID:     rr.ID,
			TTL:    -1,
		}
	}
	return resp, nil
}

func (ls *LeaseServer) LeaseLeases(ctx context.Context, rr *pb.LeaseLeasesRequest) (*pb.LeaseLeasesResponse, error) {
	resp, err := ls.le.LeaseLeases(ctx, rr)
	if err != nil && err != lease.ErrLeaseNotFound {
		return nil, togRPCError(err)
	}
	if err == lease.ErrLeaseNotFound {
		resp = &pb.LeaseLeasesResponse{
			Header: &pb.ResponseHeader{},
			Leases: []*pb.LeaseStatus{},
		}
		ls.hdr.fill(resp.Header)
	}
	return resp, nil
}

func (ls *LeaseServer) LeaseKeepAlive(stream pb.Lease_LeaseKeepAliveServer) (err error) {
	return fmt.Errorf("lease keep alive is not implemented")
}
