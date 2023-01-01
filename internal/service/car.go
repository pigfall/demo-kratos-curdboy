package service

import(
  
  "context"
  "github.com/pigfall/demo-kratos-curdboy/api"
  "github.com/pigfall/demo-kratos-curdboy/api/car/v1"
  structpb "google.golang.org/protobuf/types/known/structpb"
  "github.com/pigfall/demo-kratos-curdboy/internal/biz"
)

type CarSvc struct{
  pbcar.UnimplementedCarServer
  bizIns *biz.CarBiz
}

func NewCarSvc(bizIns *biz.CarBiz) *CarSvc{
  return &CarSvc{
    bizIns:bizIns,
  }
}

func (this *CarSvc) Create(ctx context.Context,req *structpb.Struct)(*pbcar.CarCreateResponse,error){
  data := req.AsMap()
  id,err := this.bizIns.Create(ctx,data)
  if err != nil{
    return nil,err
  }

  // TODO map from ent field type to go type.
  
    return &pbcar.CarCreateResponse{Id:int32(id)},nil
  
}

func (this *CarSvc) Query(ctx context.Context,req *common.QueryRequest) (*pbcar.CarQueryResponse,error){
  return this.bizIns.Query(ctx,req)
}
