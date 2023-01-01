package data

import(
  "context"
  "github.com/pigfall/demo-kratos-curdboy/curdboy"
  "github.com/pigfall/demo-kratos-curdboy/ent"
  "github.com/pigfall/demo-kratos-curdboy/api"
  "github.com/pigfall/demo-kratos-curdboy/internal/biz"
)

type DeptStorage struct{
  entCli *ent.Client
}

func NewDeptStorage (entCli *ent.Client) biz.DeptStorage{
  return &DeptStorage{
    entCli: entCli,
  }
}

func (this *DeptStorage) Create (ctx context.Context,data map[string]interface{}) (id int,outErr error){
  return curd.DeptCreate(ctx,data,this.entCli) 
}

func (this *DeptStorage) Query (ctx context.Context,query *common.QueryRequest) (records []*curd.Dept,count int,err error){
  q := &curd.QueryRequest{
    Filter: query.Filter,
    PageIndex: int(query.PageIndex),
    PageSize: int(query.PageSize),
  }

  records,err = curd.DeptQuery(ctx,q,this.entCli) 
  if err != nil{
    return nil,0,err
  }

  // TODO count, optimize , do not parse filter twice
  count,err = curd.DeptCount(ctx,this.entCli,q.Filter)
  if err != nil{
    return nil,0,err
  }

  return records,count,nil
}
