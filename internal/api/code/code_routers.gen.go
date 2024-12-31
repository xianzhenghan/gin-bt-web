package code

import (
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/repository/mysql"

	"go.uber.org/zap"
)

func RegisterGeneratedCodeRoutes(logger *zap.Logger, db mysql.Repo, r core.RouterGroup) {
	h := New(logger, db)

	// 新增数据
	r.POST("/code", h.Create())

	// 获取列表数据
	r.GET("/codes", h.List())

	// 根据 ID 获取数据
	r.GET("/code/:id", h.GetByID())

	// 根据 ID 更新数据
	r.PUT("/code/:id", h.UpdateByID())

	// 根据 ID 删除数据
	r.DELETE("/code/:id", h.DeleteByID())
}
