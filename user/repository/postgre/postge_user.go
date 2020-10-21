package postgre

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"lms-github/domain/v1"
)

type PgsqlUserRepository struct {
	DB *pg.DB
}

func (p PgsqlUserRepository) Update(ctx context.Context, usr *v1.User) (user *v1.User, err error) {
	_, err = p.DB.Model(usr).Update()

	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (p PgsqlUserRepository) Find(ctx context.Context, id uuid.UUID) (user *v1.User, err error) {
	user = new(v1.User)
	err = p.DB.Model(user).Where("id = ? ", id).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p PgsqlUserRepository) FindBy(ctx context.Context, key, value string) (user *v1.User, err error) {
	panic("implement me")
}

func (p PgsqlUserRepository) CreateUser(ctx context.Context, usr *v1.User) (user *v1.User, err error) {
	_, err = p.DB.Model(usr).Insert()
	if err != nil {
		return nil, err
	}

	return usr, nil

}

func (p PgsqlUserRepository) Attempt(ctx context.Context, credential *v1.Credential) (user *v1.User, err error) {
	user = new(v1.User)
	err = p.DB.Model(user).Where("email = ?", credential.Email).Select()
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewPgsqlUserRepository(db *pg.DB) v1.UserRepository {
	return PgsqlUserRepository{DB: db}
}
