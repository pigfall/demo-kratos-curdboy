package biz

import(
    "context"
    "github.com/pigfall/demo-kratos-curdboy/curdboy"
    "github.com/pigfall/demo-kratos-curdboy/api"
    "github.com/pigfall/demo-kratos-curdboy/api/user/v1"
    structpb "google.golang.org/protobuf/types/known/structpb"
    "encoding/json"
)

type UserBiz struct{
  storage UserStorage
}

func NewUserBiz (storage UserStorage) *UserBiz{
  return &UserBiz{
    storage: storage,
  }
}

type UserStorage interface{
  Create(ctx context.Context,data map[string]interface{})(id int,err error)
  Query (ctx context.Context,query *common.QueryRequest) (records []*curd.User,count int,err error)
}

func (this *UserBiz) Create (ctx context.Context, data map[string]interface{})(id int,err error){
  return this.storage.Create(ctx,data)
}

func (this *UserBiz) Query (ctx context.Context, req *common.QueryRequest) (*pbuser.UserQueryResponse, error){
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
  res := &pbuser.UserQueryResponse{
    Data: structs,
    Meta: &pbuser.UserQueryResponseMeta{
      Count:int64(count),
    },
  }

  return res,nil
}
