// do not edit, auto generated by curdboy

package curd

import(
	"context"
	"github.com/pigfall/demo-kratos-curdboy/ent"
	"github.com/pigfall/demo-kratos-curdboy/ent/predicate"
	"github.com/pigfall/curdboy/pkgs/filter"
	"github.com/pigfall/demo-kratos-curdboy/ent/dept"
	"fmt"
	"strings"
)

/*{ the number is meanning less. I just use it as mark to jump between the code
1တ   | struct Dept: as the model in db
2တ   | func (*Dept)fromEntDept: convert model in ent to curdboy node model
3တ   | func DeptCreate: func to create node
4တ   | func DeptQuery: func to query nodes by pages
5တ   | func DeptCount: func to count nodes
51တ  | func DeptUpdate: func to update nodes
52တ  | func DeptDelete: func to update nodes
6တ   | func ParseFilterToDeptPredicate: func to parse fitler string to node predicate
တ7   | functions to convert field value from interface to real type
တ8   | parse fields expr to field selector object which is used to select the field we want to query
}*/

// { 
// 1တ
type Dept struct {
  *ent.Dept
  Edges interface{} `json:"-"` // no need the edges from base class
  
}

type DeptV2 struct{
Name string
ID int

}

// 2တ
func (this *Dept) fromEntDept(base *ent.Dept,selector *DeptFieldsSelector)error{
  this.Dept = base
  return nil
}

// }

// { CURD
// 3တ 
func DeptCreate (ctx context.Context,fieldsData map[string]interface{},entCli *ent.Client)(id int,outErr error){
  mutation := entCli.Dept.Create()
  for k,v := range fieldsData {
    _ = v
    switch k {
    case "name":
      value,err := DeptNameValueFromInterface(v)
      if err != nil {
        outErr = fmt.Errorf("field name type not match for node Dept: %w",err)
        return 
      }
      mutation.SetName(value)
      case "id":
      outErr = fmt.Errorf("field id is auto generated, do not set value to the field")
      return
    default:
      outErr = fmt.Errorf("undefined field %s for node Dept",k)
      return 
    }
  }
  record,err := mutation.Save(ctx)
  if err != nil {
    outErr = err
    return
  }
  return record.ID,nil
}

// 4တ 
func DeptQuery (ctx context.Context,req *QueryRequest,entCli *ent.Client)([]*Dept,error ){
	var pred predicate.Dept
	if len(req.Filter) > 0{
    var err error
    pred,err = ParseFilterToDeptPredicate(req.Filter)
		if err != nil{
			return nil,err
		}
	}
	query := entCli.Dept.Query().Limit(req.PageSize).Offset(req.PageIndex * req.PageSize)
	if pred != nil{
		query = query.Where(pred)
	}

  // { fields selector
  fldSelector,err := ParseDeptFieldsSelector(req.Fields)
  if err != nil{
    return nil,err
  }
  // }

	records,err :=  query.All(ctx)
  if err != nil{
    return nil,err
  }
  // { convert to our node struct
  nodes := make([]*Dept,0,len(records))
  for _,r := range records{
    node := &Dept{}
    err := node.fromEntDept(r,fldSelector)
    if err != nil{
      return nil,err
    }
    nodes = append(nodes,node)
  }
  // }

  return nodes,nil
}

// 5တ 
func DeptCount (ctx context.Context,entCli *ent.Client,filter string)(int,error){
  query := entCli.Dept.Query()
  if len(filter)>0{
    pred,err := ParseFilterToDeptPredicate(filter)
    if err != nil{
      return 0, err
    }
    query.Where(pred)
  }
  return query.Count(ctx)
}

// 51တ 
func DeptUpdate (ctx context.Context,entCli *ent.Client,req *UpdateRequest)(error){
  mutation := entCli.Dept.Update()
  if len(req.Filter) > 0{
    pred,err := ParseFilterToDeptPredicate(req.Filter)
    if err != nil{
      return err
    }
    mutation.Where(pred)
  }
  for field,v := range req.Data{
    _ = v
    switch field {
    case "name":
    value,err := DeptNameValueFromInterface(v)
    if err != nil {
      err = fmt.Errorf("field name type not match for node Dept: %w",err)
      return err
    }
    mutation.SetName(value)
    case "id":
    err := fmt.Errorf("field id is auto generated, do not set value to the field")
    return err
    default:
      return fmt.Errorf("undefined field < %s > for node < Dept >",field)
    }
  }
  _,err := mutation.Save(ctx)
  return err
}

// 52တ
func DeptDelete (ctx context.Context,entCli *ent.Client,filter string)(error){
  del := entCli.Dept.Delete()
  if len(filter) > 0{
    pred,err := ParseFilterToDeptPredicate(filter)
    if err != nil{
      return err
    }
    del.Where(pred)
  }
  _,err := del.Exec(ctx)
  return err
}
// }



// { Prase filter to node predicate
// 6တ 
func ParseFilterToDeptPredicate(filterStr string)(predicate.Dept,error){
  if len(filterStr) == 0{
    return nil,fmt.Errorf("length of filter can not be 0")
  }
  filterExpr,err := filter.ParseFilter(filterStr)
  if err != nil{
    return nil,err
  }
  pred,err := ToDeptPredicate(filterExpr)
  if err != nil{
    return nil,err
  }
  return pred,nil
}



func ToDeptPredicate(filterExpr filter.Expr)(predicate.Dept,error){
visitor := &FilterVisitorDept{}
  v,err := filterExpr.Accept(visitor)
  if err != nil {
    return nil,err
  }
  return visitor.predicateFromVisitorResult(v),nil
}

type FilterVisitorDept struct{}

func (this *FilterVisitorDept) predicateFromVisitorResult(v interface{})(predicate.Dept){
	return v.(predicate.Dept)
}

func(this *FilterVisitorDept)	VisitBinaryLogicalExpr(expr *filter.BinaryLogicalExpr)(interface{},error){
	var logicalOperator = expr.Op

	left,err := expr.Left.Accept(this)
	if err != nil{
		return nil,err
	}
	right,err := expr.Right.Accept(this)
	if err != nil{
		return  nil,err
	}

	leftPred :=  this.predicateFromVisitorResult(left)
	rightPred :=  this.predicateFromVisitorResult(right)

	switch logicalOperator.Tpe{
		case filter.TokenType_KW_And:
		return dept.And(
			leftPred,
			rightPred,
		), nil
		case filter.TokenType_KW_Or:
		return dept.Or(
			leftPred,
			rightPred,
		), nil

	default:
		return nil,fmt.Errorf("unexptect logical operator %s",expr.Op.Literal)
	}
}

func(this *FilterVisitorDept)	VisitComparisionExpr(expr *filter.ComparisionExpr)(interface{},error){
	paths := strings.Split(expr.Left.GetStringValue(),".")
	if len(paths) == 1{
		var field = paths[0]
		switch field {
		
		
		case "name":
			switch expr.Op.Tpe {
				case filter.TokenType_KW_Eq:
					if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameEQ(expr.Right.GetStringValue()),nil
				case filter.TokenType_KW_Ne:
					if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameNEQ(expr.Right.GetStringValue()),nil
				case filter.TokenType_KW_Lt:
					if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameLT(expr.Right.GetStringValue()),nil
				case filter.TokenType_KW_Le:
					if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameLTE(expr.Right.GetStringValue()),nil
				case filter.TokenType_KW_Gt:
					if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameGT(expr.Right.GetStringValue()),nil
				case filter.TokenType_KW_Ge:
					if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameGTE(expr.Right.GetStringValue()),nil
				case filter.TokenType_KW_Like:if expr.Right.IsNumber(){
						return  nil,fmt.Errorf("the field name of node Dept type not matched, expect string but get number")
					}
					return dept.NameContains(expr.Right.GetStringValue()),nil
				default: 
					return nil, fmt.Errorf("unexptected comparision operator %s",expr.Op.Literal)
			}
		
		default:
			return nil, fmt.Errorf("undefined field < %s > for node < %s >",field,"Dept")
		}
	}else{ // filter for edge alias
    edge := paths[0]
    // check if edge alias exists
    switch edge{
    
    default:
      return nil,fmt.Errorf("undefined edge alis < %s > for node < Dept>",edge)
    }
	}
}

func(this *FilterVisitorDept)	VisitUnaryExpr(expr *filter.UnaryExpr)(interface{},error){
	var operator = expr.Op
	result,err := expr.Expr.Accept(this)
	if err != nil{
		return nil, err
	}
	switch operator.Tpe{
		case filter.TokenType_KW_Not:
    return dept.Not(this.predicateFromVisitorResult(result)),nil
		default:
			return nil, fmt.Errorf("unexptect unary operator %s",operator.Literal)
	}
}

// }



// {

// }

// { တ7 functions to convert field value from interface to the real type


func DeptNameValueFromInterface(v interface{})(value string,err error){
   // TODO sync with fieldTypeStr
  var ok bool
  
  
  
    value,ok = v.(string)
  

  if !ok {
    err = fmt.Errorf("The type of The field <$field.Name> of node < Dept> is  string")
    return
  }

  return
}

func DeptIDValueFromInterface(v interface{})(value int,err error){
   // TODO sync with fieldTypeStr
  var ok bool
  
  
  
    
  switch assertedV:=v.(type){
  case int:
    ok = true
    value = int(assertedV)
  case int32:
    ok = true
    value = int(assertedV)
  case int64:
    ok = true
    value = int(assertedV)
  case float32:
    ok = true
    value = int(assertedV)
  case float64:
    ok = true
    value = int(assertedV)
  }

  

  if !ok {
    err = fmt.Errorf("The type of The field <$field.Name> of node < Dept> is  int")
    return
  }

  return
}

// }






// { တ8 parse fields expr to field selector object which is used to select the field we want to query
func ParseDeptFieldsSelector(fieldsStr string)(*DeptFieldsSelector,error){
  fields := strings.Split(fieldsStr,",")
  selector := &DeptFieldsSelector{
    Fields: make([]string,0,len(fields)),
  }
  for _,f := range fields {
    err := selector.AddField(f)
    if err != nil{
      return nil, err
    }
  }
  return selector,nil
}

type DeptFieldsSelector struct{
  SelectAllField bool
  Fields []string
}

func (this *DeptFieldsSelector) AddField(field string) error{
  elems := strings.Split(field,".")
  if  len(elems) == 1{
    f := elems[0]
    if f == "*"{
      this.SelectAllField = true
      return nil
    }
    this.Fields = append(this.Fields,f)
  }else{
      edgeAlias := elems[0]
      switch edgeAlias{
        default:
          return fmt.Errorf("undefined edge alias %s for node Dept",edgeAlias)
      }
  }
  return nil
}

func (this *DeptFieldsSelector) Select(query *ent.DeptQuery)error{
  if !this.SelectAllField{
    query.Select(this.Fields...)
  }
  return nil
}

// }
