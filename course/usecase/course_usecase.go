package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	v1 "lms-github/domain/v1"
	"net/http"
	"time"
)

type CourseUsecase struct {
	CourseRepo     v1.CourseRepository
	ContextTimeout time.Duration
}

func NewCourseUsecase(CourseRepo v1.CourseRepository, timeout time.Duration) v1.CourseUsecase {
	return CourseUsecase{
		CourseRepo:     CourseRepo,
		ContextTimeout: timeout,
	}
}

func (c CourseUsecase) CreateCourse(ctx context.Context, crs *v1.Course, form *http.Request) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, c.ContextTimeout)
	defer cancel()

	//assign to struct
	crs.ID = uuid.New()
	crs.Title = form.FormValue("title")
	crs.Description = form.FormValue("description")
	crs.CreatedAt = time.Now()
	crs.Slug = slug.Make(crs.Title)
	crs.Thumbnail = crs.Title

	course, err := c.CourseRepo.CreateCourse(ctx, crs)

	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CourseUsecase) GetCourseBySlug(ctx context.Context, slug string) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, c.ContextTimeout)
	defer cancel()
	course, err := c.CourseRepo.FindBy(ctx, "slug", slug)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CourseUsecase) EnrollCourse(ctx context.Context, userID, courseID uuid.UUID) (res interface{}, err error) {
	panic("implement me")
}

func (c CourseUsecase) DeleteCourse(ctx context.Context, id uuid.UUID) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, c.ContextTimeout)
	defer cancel()

	err = c.CourseRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
