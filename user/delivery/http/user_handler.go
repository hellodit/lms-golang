package http

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	v1 "lms-github/domain/v1"
	"lms-github/middleware"
	"net/http"
)

type userHandler struct {
	userUsecase v1.UserUsecase
}

func NewUserHandler(e *echo.Echo, UserUsecase v1.UserUsecase) {
	handler := &userHandler{userUsecase: UserUsecase}

	user := e.Group("/user")
	customMiddleware := middleware.Init()
	user.GET("/:id", handler.GetByIDHandler, customMiddleware.Auth)
	user.POST("/register", handler.RegisterHandler)
	user.POST("/login", handler.LoginHandler)
	user.POST("/update", handler.UpdateHandler)
}

func (u userHandler) UpdateHandler(e echo.Context) error {
	panic("implement me")
}

func (u userHandler) GetByIDHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}
	id, err := uuid.Parse(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}
	res, err := u.userUsecase.GetUserById(ctx, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

func (u userHandler) RegisterHandler(e echo.Context) error {
	rules := govalidator.MapData{
		"name":     []string{"required"},
		"password": []string{"required"},
		"email":    []string{"required"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	ctx := e.Request().Context()
	var usr v1.User

	if ctx == nil {
		ctx = context.Background()
	}

	res, err := u.userUsecase.Register(ctx, &usr, e.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

func (u userHandler) LoginHandler(e echo.Context) error {
	rules := govalidator.MapData{
		"password": []string{"required"},
		"email":    []string{"required"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	var credential v1.Credential

	if err := e.Bind(&credential); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(errors.New("invalid parameter"))
	}

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := u.userUsecase.Login(ctx, &credential)

	if err != nil {
		return echo.NewHTTPError(http.StatusFailedDependency, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, res)
}
