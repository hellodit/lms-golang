package postgre

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	v1 "lms-github/domain/v1"
)

type PgsqlCourseRepository struct {
	DB *pg.DB
}

func (p PgsqlCourseRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	course := new(v1.Course)
	_, err = p.DB.Model(course).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (p PgsqlCourseRepository) CreateCourse(ctx context.Context, crs *v1.Course) (course *v1.Course, err error) {
	_, err = p.DB.Model(crs).Insert()
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func (p PgsqlCourseRepository) Update(ctx context.Context, crs *v1.Course) (course *v1.Course, err error) {
	_, err = p.DB.Model(crs).Update()

	if err != nil {
		return nil, err
	}

	return crs, nil
}

func (p PgsqlCourseRepository) Find(ctx context.Context, id uuid.UUID) (course *v1.Course, err error) {
	course = new(v1.Course)
	err = p.DB.Model(course).Where("id = ? ", id).First()
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (p PgsqlCourseRepository) FindBy(ctx context.Context, key, value string) (course *v1.Course, err error) {
	course = new(v1.Course)
	err = p.DB.Model(course).Where(key+"= ? ", value).First()
	if err != nil {
		return nil, err
	}
	return course, nil
}

func NewPgsqlCourseRepository(db *pg.DB) v1.CourseRepository {
	return PgsqlCourseRepository{DB: db}
}
