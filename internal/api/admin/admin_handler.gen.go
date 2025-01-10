package admin

import (
	"bt-web-ide/configs"
	"bt-web-ide/internal/pkg/jwtoken"
	"bt-web-ide/internal/pkg/password"
	"bt-web-ide/internal/proposal"
	"net/http"
	"strconv"
	"time"

	"bt-web-ide/internal/code"
	"bt-web-ide/internal/pkg/core"
	"bt-web-ide/internal/repository/mysql"
	"bt-web-ide/internal/repository/mysql/dao"
	"bt-web-ide/internal/repository/mysql/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type handler struct {
	logger  *zap.Logger
	writeDB *dao.Query
	readDB  *dao.Query
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RegisterResponse struct {
	Token     string   `json:"token"`
	ExpireAt  int64    `json:"expire_at"`
	TokenType string   `json:"token_type"`
	User      Register `json:"user"`
}

type genResultInfo struct {
	RowsAffected int64 `json:"rows_affected"`
	Error        error `json:"error"`
}

func New(logger *zap.Logger, db mysql.Repo) *handler {
	return &handler{
		logger:  logger,
		writeDB: dao.Use(db.GetDbW()),
		readDB:  dao.Use(db.GetDbR()),
	}
}

// Create 新增数据
// @Summary 新增数据
// @Description 新增数据
// @Tags Table.admin
// @Accept json
// @Produce json
// @Param RequestBody body model.Admin true "请求参数"
// @Success 200 {object} model.Admin
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		var createData model.Admin
		if err := ctx.ShouldBindJSON(&createData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}

		if err := h.writeDB.Admin.WithContext(ctx.RequestContext()).Create(&createData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ServerError,
				err.Error()),
			)
			return
		}

		ctx.Payload(createData)
	}
}

// List 获取列表数据
// @Summary 获取列表数据
// @Description 获取列表数据
// @Tags Table.admin
// @Accept json
// @Produce json
// @Success 200 {object} []model.Admin
// @Failure 400 {object} code.Failure
// @Router /api/admins [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		list, err := h.readDB.Admin.WithContext(ctx.RequestContext()).Find()
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ServerError,
				err.Error()),
			)
			return
		}

		ctx.Payload(list)
	}
}

// GetByID 根据 ID 获取数据
// @Summary 根据 ID 获取数据
// @Description 根据 ID 获取数据
// @Tags Table.admin
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Admin
// @Failure 400 {object} code.Failure
// @Router /api/admin/{id} [get]
func (h *handler) GetByID() core.HandlerFunc {
	return func(ctx core.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}

		info, err := h.readDB.Admin.WithContext(ctx.RequestContext()).Where(h.readDB.Admin.ID.Eq(int32(id))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ServerError,
					"record not found"),
				)
			} else {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ServerError,
					err.Error()),
				)
			}
			return
		}

		ctx.Payload(info)
	}
}

// DeleteByID 根据 ID 删除数据
// @Summary 根据 ID 删除数据
// @Description 根据 ID 删除数据
// @Tags Table.admin
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} genResultInfo
// @Failure 400 {object} code.Failure
// @Router /api/admin/{id} [delete]
func (h *handler) DeleteByID() core.HandlerFunc {
	return func(ctx core.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}

		info, err := h.readDB.Admin.WithContext(ctx.RequestContext()).Where(h.readDB.Admin.ID.Eq(int32(id))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ServerError,
					"record not found"),
				)
			} else {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ServerError,
					err.Error()),
				)
			}
			return
		}

		result, err := h.writeDB.Admin.Delete(info)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ServerError,
				err.Error()),
			)
		}

		resultInfo := new(genResultInfo)
		resultInfo.RowsAffected = result.RowsAffected
		resultInfo.Error = result.Error

		ctx.Payload(resultInfo)
	}
}

// UpdateByID 根据 ID 更新数据
// @Summary 根据 ID 更新数据
// @Description 根据 ID 更新数据
// @Tags Table.admin
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param RequestBody body model.Admin true "请求参数"
// @Success 200 {object} genResultInfo
// @Failure 400 {object} code.Failure
// @Router /api/admin/{id} [put]
func (h *handler) UpdateByID() core.HandlerFunc {
	return func(ctx core.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}

		var updateData map[string]interface{}
		if err := ctx.ShouldBindJSON(&updateData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ServerError,
				err.Error()),
			)
			return
		}

		info, err := h.readDB.Admin.WithContext(ctx.RequestContext()).Where(h.readDB.Admin.ID.Eq(int32(id))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ServerError,
					"record not found"),
				)
			} else {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ServerError,
					err.Error()),
				)
			}
			return
		}

		result, err := h.writeDB.Admin.WithContext(ctx.RequestContext()).Where(h.writeDB.Admin.ID.Eq(info.ID)).Updates(updateData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ServerError,
				err.Error()),
			)
			return
		}

		resultInfo := new(genResultInfo)
		resultInfo.RowsAffected = result.RowsAffected
		resultInfo.Error = result.Error

		ctx.Payload(resultInfo)
	}
}

func (h *handler) Register() core.HandlerFunc {
	return func(ctx core.Context) {
		var registerData Register
		if err := ctx.ShouldBindJSON(&registerData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}
		if registerData.Username == "" || registerData.Password == "" {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				"password or username or password is empty",
			))
			return
		}
		var createData model.UsersRegistration
		createData.Username = registerData.Username
		createData.Email = registerData.Email
		token := password.GenerateLoginToken(registerData.Username)
		createData.PasswordHash = token
		createData.PhoneNumber = registerData.Phone
		if err := h.writeDB.UsersRegistration.WithContext(ctx.RequestContext()).Create(&createData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ServerError,
				err.Error()),
			)
			return
		}
		// return token

		sessionUserInfo := proposal.SessionUserInfo{
			UserName: registerData.Username,
			PassWord: registerData.Password,
		}
		tokenString, err := jwtoken.New(configs.Get().JWT.Secret).Sign(sessionUserInfo, 24*time.Hour)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}
		r := RegisterResponse{
			Token:     tokenString,
			ExpireAt:  time.Now().Add(time.Hour * 24).Unix(),
			TokenType: "Bearer",
			User: Register{
				Username: registerData.Username,
				Password: registerData.Password,
				Email:    registerData.Email,
			},
		}
		ctx.PayloadWithCode(r, "00000", "success")
	}
}

func (h *handler) Login() core.HandlerFunc {
	return func(ctx core.Context) {
		var loginData Login
		if err := ctx.ShouldBindJSON(&loginData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}
		if loginData.Username == "" || loginData.Password == "" {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				"password or username or password is empty",
			))
			return
		}

		sessionUserInfo := proposal.SessionUserInfo{
			UserName: loginData.Username,
			PassWord: loginData.Password,
		}
		tokenString, err := jwtoken.New(configs.Get().JWT.Secret).Sign(sessionUserInfo, 24*time.Hour)

		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				err.Error()),
			)
			return
		}
		ctx.PayloadWithCode(tokenString, "00000", "success")
	}
}
