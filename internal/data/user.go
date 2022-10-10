package data

import(
  "context"
  "github.com/pigfall/demo-kratos-curdboy/curdboy"
  "github.com/pigfall/demo-kratos-curdboy/ent"
  "github.com/pigfall/demo-kratos-curdboy/api"
  "github.com/pigfall/demo-kratos-curdboy/internal/biz"
)

type UserStorage struct{
  entCli *ent.Client
}

func NewUserStorage (entCli *ent.Client) biz.UserStorage{
  return &UserStorage{
    entCli: entCli,
  }
}

func (this *UserStorage) Create (ctx context.Context,data map[string]interface{}) (id int,outErr error){
  return curd.UserCreate(ctx,data,this.entCli) 
}

func (this *UserStorage) Query (ctx context.Context,query *common.QueryRequest) (records []*curd.User,count int,err error){
  q := &curd.QueryRequest{
    Filter: query.Filter,
    PageIndex: int(query.PageIndex),
    PageSize: int(query.PageSize),
  }

  records,err = curd.UserQuery(ctx,q,this.entCli) 
  if err != nil{
    return nil,0,err
  }

  // TODO count, optimize , do not parse filter twice
  count,err = curd.UserCount(ctx,this.entCli,q.Filter)
  if err != nil{
    return nil,0,err
  }

  return records,count,nil
}
