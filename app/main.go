package main

import (
	"lms-github/config"
	"lms-github/db/postgre"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	_courseHttpDelivery "lms-github/course/delivery/http"
	_userHttpDelivery "lms-github/user/delivery/http"

	_coursePostgresRepository "lms-github/course/repository/postgre"
	_userPostgreRepository "lms-github/user/repository/postgre"

	_courseUsecase "lms-github/course/usecase"
	_userUsecase "lms-github/user/usecase"
)

func main() {
	config.ReadConfig()
	db := postgre.Connect()

	timeoutCtx := time.Duration(5) * time.Second

	server := &http.Server{
		Addr:         ":" + viper.GetString("server.port"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world from "+viper.GetString("app_name"))
	})

	//Call repositories
	userRepo := _userPostgreRepository.NewPgsqlUserRepository(db)
	courseRepo := _coursePostgresRepository.NewPgsqlCourseRepository(db)

	//call usecase
	useUsecae := _userUsecase.NewUserUseCase(userRepo, timeoutCtx)
	courseUsecase := _courseUsecase.NewCourseUsecase(courseRepo, timeoutCtx)

	//call delivery
	_userHttpDelivery.NewUserHandler(e, useUsecae)
	_courseHttpDelivery.NewCourseHandler(e, courseUsecase)

	err := e.StartServer(server)
	if err != nil {
		e.Logger.Info("Shutting down the server")
	}

}
