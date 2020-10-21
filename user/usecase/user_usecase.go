package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	v1 "lms-github/domain/v1"
	"lms-github/helper"
	"net/http"
	"time"
)

type userUsecase struct {
	UserRepo       v1.UserRepository
	ContextTimeout time.Duration
}

func (u userUsecase) UpdateUser(ctx context.Context, usr *v1.User, form *http.Request) (res interface{}, err error) {
	panic("implement me")
}

func (u userUsecase) GetUserById(ctx context.Context, id uuid.UUID) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	user, err := u.UserRepo.Find(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUsecase) Register(ctx context.Context, usr *v1.User, form *http.Request) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.FormValue("password")), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	usr.ID = uuid.New()
	usr.Name = form.FormValue("name")
	usr.Email = form.FormValue("email")
	usr.CreatedAt = time.Now()
	usr.Password = string(hashedPassword)

	user, err := u.UserRepo.CreateUser(ctx, usr)

	if err != nil {
		return
	}

	return user, nil
}

func (u userUsecase) Login(ctx context.Context, credential *v1.Credential) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	user, err := u.UserRepo.Attempt(ctx, credential)
	if err != nil {
		return nil, err
	}

	token, exp, err := helper.GenerateJwt(ctx, user)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token_type":   "Bearer",
		"access_token": token,
		"expires_in":   exp,
		"profile":      user,
	}, nil

}

func (u userUsecase) Logout(ctx context.Context, claims jwt.Claims) {
	panic("implement me")
}

func NewUserUseCase(UserRepo v1.UserRepository, timeout time.Duration) v1.UserUsecase {
	return userUsecase{
		UserRepo:       UserRepo,
		ContextTimeout: timeout,
	}
}
