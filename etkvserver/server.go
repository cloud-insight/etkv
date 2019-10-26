package etkvserver

import (
	"context"
	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
	"github.com/tikv/client-go/txnkv"
	"go.etcd.io/etcd/lease"
	"go.uber.org/zap"

	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
)

type Authenticator interface {
	AuthEnable(ctx context.Context, r *pb.AuthEnableRequest) (*pb.AuthEnableResponse, error)
	AuthDisable(ctx context.Context, r *pb.AuthDisableRequest) (*pb.AuthDisableResponse, error)
	Authenticate(ctx context.Context, r *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error)
	UserAdd(ctx context.Context, r *pb.AuthUserAddRequest) (*pb.AuthUserAddResponse, error)
	UserDelete(ctx context.Context, r *pb.AuthUserDeleteRequest) (*pb.AuthUserDeleteResponse, error)
	UserChangePassword(ctx context.Context, r *pb.AuthUserChangePasswordRequest) (*pb.AuthUserChangePasswordResponse, error)
	UserGrantRole(ctx context.Context, r *pb.AuthUserGrantRoleRequest) (*pb.AuthUserGrantRoleResponse, error)
	UserGet(ctx context.Context, r *pb.AuthUserGetRequest) (*pb.AuthUserGetResponse, error)
	UserRevokeRole(ctx context.Context, r *pb.AuthUserRevokeRoleRequest) (*pb.AuthUserRevokeRoleResponse, error)
	RoleAdd(ctx context.Context, r *pb.AuthRoleAddRequest) (*pb.AuthRoleAddResponse, error)
	RoleGrantPermission(ctx context.Context, r *pb.AuthRoleGrantPermissionRequest) (*pb.AuthRoleGrantPermissionResponse, error)
	RoleGet(ctx context.Context, r *pb.AuthRoleGetRequest) (*pb.AuthRoleGetResponse, error)
	RoleRevokePermission(ctx context.Context, r *pb.AuthRoleRevokePermissionRequest) (*pb.AuthRoleRevokePermissionResponse, error)
	RoleDelete(ctx context.Context, r *pb.AuthRoleDeleteRequest) (*pb.AuthRoleDeleteResponse, error)
	UserList(ctx context.Context, r *pb.AuthUserListRequest) (*pb.AuthUserListResponse, error)
	RoleList(ctx context.Context, r *pb.AuthRoleListRequest) (*pb.AuthRoleListResponse, error)
}

type ServerConfig struct {
	MaxTxnOps uint
	Logger *zap.Logger
}

type EtkvCluster struct {
	id int64
}

func (ec *EtkvCluster) ID() int64 {
	return ec.id
}

type EtkvServer struct {
	id       int64
	ec       *EtkvCluster
	Cfg      ServerConfig

	rawKvClient *rawkv.Client
	txnKvClient *txnkv.Client
}

func NewEtkvServer(ctx context.Context, pdAddrs []string, conf config.Config) (*EtkvServer, error) {
	rawKvClient, err := rawkv.NewClient(ctx, pdAddrs, conf)
	if err != nil {
		return nil, err
	}

	txnKvClient, err := txnkv.NewClient(ctx, pdAddrs, conf)
	if err != nil {
		return nil, err
	}

	etkvServer := &EtkvServer{
		rawKvClient: rawKvClient,
		txnKvClient: txnKvClient,
	}
	return etkvServer, nil
}

func (es *EtkvServer) ID() int64 {
	return es.id
}

func (es *EtkvServer) Cluster() *EtkvCluster {
	return es.ec
}

func (es *EtkvServer) Range(ctx context.Context, r *pb.RangeRequest) (*pb.RangeResponse, error) {
	return nil, nil
}

func (es *EtkvServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	return nil, nil
}

func (es *EtkvServer) DeleteRange(ctx context.Context, r *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	return nil, nil
}

func (es *EtkvServer) Txn(ctx context.Context, r *pb.TxnRequest) (*pb.TxnResponse, error) {
	return nil, nil
}

func (es *EtkvServer) Compact(ctx context.Context, r *pb.CompactionRequest) (*pb.CompactionResponse, error) {
	return nil, nil
}

func (es EtkvServer) LeaseGrant(ctx context.Context, r *pb.LeaseGrantRequest) (*pb.LeaseGrantResponse, error) {
	return nil, nil
}

func (es EtkvServer) LeaseRevoke(ctx context.Context, r *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	return nil, nil
}

func (es EtkvServer) LeaseRenew(ctx context.Context, id lease.LeaseID) (int64, error) {
	return 0, nil
}

func (es EtkvServer) LeaseTimeToLive(ctx context.Context, r *pb.LeaseTimeToLiveRequest) (*pb.LeaseTimeToLiveResponse, error) {
	return nil, nil
}

func (es EtkvServer) LeaseLeases(ctx context.Context, r *pb.LeaseLeasesRequest) (*pb.LeaseLeasesResponse, error) {
	return nil, nil
}
