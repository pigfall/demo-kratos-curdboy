package server

import (
	"github.com/gorilla/handlers"
	v1Car "github.com/pigfall/demo-kratos-curdboy/api/car/v1"
	v1Dept "github.com/pigfall/demo-kratos-curdboy/api/dept/v1"
	v1 "github.com/pigfall/demo-kratos-curdboy/api/user/v1"
	"github.com/pigfall/demo-kratos-curdboy/internal/conf"
	"github.com/pigfall/demo-kratos-curdboy/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, userSvc *service.UserSvc, deptSvc *service.DeptSvc, carSvc *service.CarSvc, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(
		opts,
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedHeaders([]string{"Content-Type"}),
				//handlers.AllowedOrigins([]string{"*"}),
			),
		),
	)
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, userSvc)
	v1Dept.RegisterDeptHTTPServer(srv, deptSvc)
	v1Car.RegisterCarHTTPServer(srv, carSvc)
	return srv
}
