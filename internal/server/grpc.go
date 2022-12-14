package server

import (
	v1Car "github.com/pigfall/demo-kratos-curdboy/api/car/v1"
	v1Dept "github.com/pigfall/demo-kratos-curdboy/api/dept/v1"
	v1 "github.com/pigfall/demo-kratos-curdboy/api/user/v1"
	"github.com/pigfall/demo-kratos-curdboy/internal/conf"
	"github.com/pigfall/demo-kratos-curdboy/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, userSvc *service.UserSvc, deptSvc *service.DeptSvc, carSvc *service.CarSvc, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, userSvc)
	v1Dept.RegisterDeptServer(srv, deptSvc)
	v1Car.RegisterCarServer(srv, carSvc)
	return srv
}
