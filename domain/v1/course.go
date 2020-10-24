package v1

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type (
	Course struct {
		tableName   struct{}  `pg:"courses"`
		ID          uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
		Title       string    `json:"title" form:"title"`
		Slug        string    `json:"slug" form:"slug"`
		Description string    `json:"description" form:"description"`
		Thumbnail   string    `json:"thumbnail" form:"thumbnail"`
		CreatedAt   time.Time `pg:"default:now()" json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		//Users		[]User 		`pg:"many2many:user_course"`
		//Topic 		[]Topic		`pg:"rel:has-many"`
	}
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, crs *Course) (course *Course, err error)
	Update(ctx context.Context, crs *Course) (course *Course, err error)
	Find(ctx context.Context, id uuid.UUID) (course *Course, err error)
	FindBy(ctx context.Context, key, value string) (course *Course, err error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type CourseUsecase interface {
	CreateCourse(ctx context.Context, crs *Course, form *http.Request) (res interface{}, err error)
	GetCourseBySlug(ctx context.Context, slug string) (res interface{}, err error)
	EnrollCourse(ctx context.Context, userID, courseID uuid.UUID) (res interface{}, err error)
	DeleteCourse(ctx context.Context, id uuid.UUID) (res interface{}, err error)
}
