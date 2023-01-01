package service

import (
	"context"

	common "github.com/pigfall/demo-kratos-curdboy/api"
	pbuser "github.com/pigfall/demo-kratos-curdboy/api/user/v1"
	"github.com/pigfall/demo-kratos-curdboy/internal/biz"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

type UserSvc struct {
	pbuser.UnimplementedUserServer
	bizIns *biz.UserBiz
}

func NewUserSvc(bizIns *biz.UserBiz) *UserSvc {
	return &UserSvc{
		bizIns: bizIns,
	}
}

func (this *UserSvc) Create(ctx context.Context, req *structpb.Struct) (*pbuser.UserCreateResponse, error) {
	data := req.AsMap()
	id, err := this.bizIns.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	// TODO map from ent field type to go type.

	return &pbuser.UserCreateResponse{Id: int32(id)}, nil

}

func (this *UserSvc) Query(ctx context.Context, req *common.QueryRequest) (*pbuser.UserQueryResponse, error) {
	return this.bizIns.Query(ctx, req)
}
