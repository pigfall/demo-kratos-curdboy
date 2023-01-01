package data

import(
  "context"
  "github.com/pigfall/demo-kratos-curdboy/curdboy"
  "github.com/pigfall/demo-kratos-curdboy/ent"
  "github.com/pigfall/demo-kratos-curdboy/api"
  "github.com/pigfall/demo-kratos-curdboy/internal/biz"
)

type CarStorage struct{
  entCli *ent.Client
}

func NewCarStorage (entCli *ent.Client) biz.CarStorage{
  return &CarStorage{
    entCli: entCli,
  }
}

func (this *CarStorage) Create (ctx context.Context,data map[string]interface{}) (id int,outErr error){
  return curd.CarCreate(ctx,data,this.entCli) 
}

func (this *CarStorage) Query (ctx context.Context,query *common.QueryRequest) (records []*curd.Car,count int,err error){
  q := &curd.QueryRequest{
    Filter: query.Filter,
    PageIndex: int(query.PageIndex),
    PageSize: int(query.PageSize),
  }

  records,err = curd.CarQuery(ctx,q,this.entCli) 
  if err != nil{
    return nil,0,err
  }

  // TODO count, optimize , do not parse filter twice
  count,err = curd.CarCount(ctx,this.entCli,q.Filter)
  if err != nil{
    return nil,0,err
  }

  return records,count,nil
}
