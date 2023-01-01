package biz

import(
    "context"
    "github.com/pigfall/demo-kratos-curdboy/curdboy"
    "github.com/pigfall/demo-kratos-curdboy/api"
    "github.com/pigfall/demo-kratos-curdboy/api/dept/v1"
    structpb "google.golang.org/protobuf/types/known/structpb"
    "encoding/json"
)

type DeptBiz struct{
  storage DeptStorage
}

func NewDeptBiz (storage DeptStorage) *DeptBiz{
  return &DeptBiz{
    storage: storage,
  }
}

type DeptStorage interface{
  Create(ctx context.Context,data map[string]interface{})(id int,err error)
  Query (ctx context.Context,query *common.QueryRequest) (records []*curd.Dept,count int,err error)
}

func (this *DeptBiz) Create (ctx context.Context, data map[string]interface{})(id int,err error){
  return this.storage.Create(ctx,data)
}

func (this *DeptBiz) Query (ctx context.Context, req *common.QueryRequest) (*pbdept.DeptQueryResponse, error){
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
  res := &pbdept.DeptQueryResponse{
    Data: structs,
    Meta: &pbdept.DeptQueryResponseMeta{
      Count:int64(count),
    },
  }

  return res,nil
}
