package service

import(
  
  "context"
  "github.com/pigfall/demo-kratos-curdboy/api"
  "github.com/pigfall/demo-kratos-curdboy/api/user/v1"
  structpb "google.golang.org/protobuf/types/known/structpb"
  "github.com/pigfall/demo-kratos-curdboy/internal/biz"
)

type UserSvc struct{
  pbuser.UnimplementedUserServer
  bizIns *biz.UserBiz
}

func NewUserSvc(bizIns *biz.UserBiz) *UserSvc{
  return &UserSvc{
    bizIns:bizIns,
  }
}

func (this *UserSvc) Create(ctx context.Context,req *structpb.Struct)(id int,err error){
  data := req.AsMap()
  return this.bizIns.Create(ctx,data)
}

func (this *UserSvc) Query(ctx context.Context,req *common.QueryRequest) (*pbuser.UserQueryResponse,error){
  return this.bizIns.Query(ctx,req)
}
