package biz

import(
    "context"
    "github.com/pigfall/demo-kratos-curdboy/curdboy"
    "github.com/pigfall/demo-kratos-curdboy/api"
    "github.com/pigfall/demo-kratos-curdboy/api/car/v1"
    structpb "google.golang.org/protobuf/types/known/structpb"
    "encoding/json"
)

type CarBiz struct{
  storage CarStorage
}

func NewCarBiz (storage CarStorage) *CarBiz{
  return &CarBiz{
    storage: storage,
  }
}

type CarStorage interface{
  Create(ctx context.Context,data map[string]interface{})(id int,err error)
  Query (ctx context.Context,query *common.QueryRequest) (records []*curd.Car,count int,err error)
}

func (this *CarBiz) Create (ctx context.Context, data map[string]interface{})(id int,err error){
  return this.storage.Create(ctx,data)
}

func (this *CarBiz) Query (ctx context.Context, req *common.QueryRequest) (*pbcar.CarQueryResponse, error){
  records,count,err := this.storage.Query(ctx,req)
  if err != nil{
    return nil, err
  }
  // { record to structpb.Struct
  structs := make([]*structpb.Struct,0,len(records))
  for _,record := range records{
    s := &structpb.Struct{}
    bytes,err := json.Marshal(record)
    if err != nil{
      return nil,err
    }
    err = s.UnmarshalJSON(bytes)
    if err != nil{
      return nil,err
    }
    structs = append(structs,s)
  }

  // }
  res := &pbcar.CarQueryResponse{
    Data: structs,
    Meta: &pbcar.CarQueryResponseMeta{
      Count:int64(count),
    },
  }

  return res,nil
}
