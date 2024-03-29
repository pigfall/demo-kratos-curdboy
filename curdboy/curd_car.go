// do not edit, auto generated by curdboy

package curd

import(
	"context"
	"github.com/pigfall/demo-kratos-curdboy/ent"
	"github.com/pigfall/demo-kratos-curdboy/ent/predicate"
	"github.com/pigfall/curdboy/pkgs/filter"
	"github.com/pigfall/demo-kratos-curdboy/ent/car"
	"fmt"
	"strings"
)

/*{ the number is meanning less. I just use it as mark to jump between the code
1တ   | struct Car: as the model in db
2တ   | func (*Car)fromEntCar: convert model in ent to curdboy node model
3တ   | func CarCreate: func to create node
4တ   | func CarQuery: func to query nodes by pages
5တ   | func CarCount: func to count nodes
51တ  | func CarUpdate: func to update nodes
52တ  | func CarDelete: func to update nodes
6တ   | func ParseFilterToCarPredicate: func to parse fitler string to node predicate
တ7   | functions to convert field value from interface to real type
တ8   | parse fields expr to field selector object which is used to select the field we want to query
}*/

// { 
// 1တ
type Car struct {
  *ent.Car
  Edges interface{} `json:"-"` // no need the edges from base class
  
}

type CarV2 struct{
Name string
ID int
Owner 