//go:build tools

package tools

import (
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2"
	_ "github.com/google/gnostic/cmd/protoc-gen-openapi"
	_ "github.com/pigfall/go-kratos-curdboy/cmd/kratos-cbc"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
