package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	v1 "lms-github/domain/v1"
	"lms-github/middleware"
	"net/http"
)

type courseHanlder struct {
	courseUsecase v1.CourseUsecase
}

func NewCourseHandler(e *echo.Echo, CourseUsecase v1.CourseUsecase) {
	handler := &courseHanlder{courseUsecase: CourseUsecase}
	customMiddleware := middleware.Init()
	course := e.Group("/course")
	course.Use(customMiddleware.Auth)
	course.GET("/:slug", handler.GetCourseHandler)
	course.POST("/store", handler.CreateCourseHandler)
}

func (c courseHanlder) CreateCourseHandler(e echo.Context) error {

	rules := govalidator.MapData{
		"title":       []string{"required"},
		"description": []string{"required"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}
	ctx := e.Request().Context()
	var crs v1.Course

	if ctx == nil {
		ctx = context.Background()
	}

	res, err := c.courseUsecase.CreateCourse(ctx, &crs, e.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})

}

func (c courseHanlder) GetCourseHandler(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	slugString := e.Param("slug")
	fmt.Print(slugString)
	res, err := c.courseUsecase.GetCourseBySlug(ctx, slugString)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}
