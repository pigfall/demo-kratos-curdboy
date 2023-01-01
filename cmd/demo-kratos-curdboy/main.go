package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "github.com/go-sql-driver/mysql"

	"github.com/pigfall/demo-kratos-curdboy/ent"
	"github.com/pigfall/demo-kratos-curdboy/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	logHelper := log.NewHelper(logger)
	logHelper.Info("connecting to db")
	drv, err := sql.Open(bc.GetData().GetDatabase().Driver, bc.GetData().GetDatabase().Source)
	if err != nil {
		panic(err)
	}
	if err := drv.DB().Ping(); err != nil {
		panic(err)
	}
	entClient := ent.NewClient(ent.Driver(drv))
	logHelper.Info("db connected")
	if err := entClient.Schema.Create(context.Background()); err != nil {
		err = fmt.Errorf("migrate schema error: %w\n", err)
		logHelper.Error(err)
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger, entClient)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
