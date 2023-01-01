package service

import(
  
  "context"
  "github.com/pigfall/demo-kratos-curdboy/api"
  "github.com/pigfall/demo-kratos-curdboy/api/dept/v1"
  structpb "google.golang.org/protobuf/types/known/structpb"
  "github.com/pigfall/demo-kratos-curdboy/internal/biz"
)

type DeptSvc struct{
  pbdept.UnimplementedDeptServer
  bizIns *biz.DeptBiz
}

func NewDeptSvc(bizIns *biz.DeptBiz) *DeptSvc{
  return &DeptSvc{
    bizIns:bizIns,
  }
}

func (this *DeptSvc) Create(ctx context.Context,req *structpb.Struct)(*pbdept.DeptCreateResponse,error){
  data := req.AsMap()
  id,err := this.bizIns.Create(ctx,data)
  if err != nil{
    return nil,err
  }

  // TODO map from ent field type to go type.
  
    return &pbdept.DeptCreateResponse{Id:int32(id)},nil
  
}

func (this *DeptSvc) Query(ctx context.Context,req *common.QueryRequest) (*pbdept.DeptQueryResponse,error){
  return this.bizIns.Query(ctx,req)
}
