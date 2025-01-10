package router

import (
	"bt-web-ide/internal/api/admin"
	"bt-web-ide/internal/api/code"
	"bt-web-ide/internal/pkg/core"
	"bt-web-ide/internal/repository/mysql"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewHTTPMux(logger *zap.Logger, db mysql.Repo) (core.Mux, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	if db == nil {
		return nil, errors.New("db required")
	}

	mux, err := core.New(logger,
		core.WithEnableCors(),
		core.WithEnableSwagger(),
		core.WithEnablePProf(),
	)

	if err != nil {
		panic(err)
	}

	// 定义自动生成的路由组前缀为 /api
	generatedRouterGroup := mux.Group("/api")

	// 注册路由
	admin.RegisterGeneratedAdminRoutes(logger, db, generatedRouterGroup)
	code.RegisterGeneratedCodeRoutes(logger, db, generatedRouterGroup)
	return mux, nil
}
