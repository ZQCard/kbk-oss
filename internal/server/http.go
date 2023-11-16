package server

import (
	ossV1 "repo.example.com/oss/api/oss/v1"
	"repo.example.com/oss/internal/conf"
	"repo.example.com/oss/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(conf *conf.Bootstrap, server *conf.Server, service *service.FileService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			// 元信息
			metadata.Server(),
			// 访问日志
			logging.Server(logger),
		),
	}
	if server.Http.Network != "" {
		opts = append(opts, http.Network(server.Http.Network))
	}
	if server.Http.Addr != "" {
		opts = append(opts, http.Address(server.Http.Addr))
	}
	if server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	ossV1.RegisterOSSHTTPServer(srv, service)
	return srv
}
