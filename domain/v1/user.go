package v1

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type (
	Credential struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	User struct {
		tableName struct{}  `pg:"users"`
		ID        uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
		Name      string    `pg:"name,type:varchar(255)" json:"name" form:"name"`
		Email     string    `pg:"email,type:varchar(255)" json:"email" form:"email"`
		Password  string    `pg:"password,type:varchar(255)" json:"-" form:"password"`
		CreatedAt time.Time `pg:"default:now()" json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

type UserRepository interface {
	CreateUser(ctx context.Context, usr *User) (user *User, err error)
	Attempt(ctx context.Context, credential *Credential) (user *User, err error)
	Update(ctx context.Context, usr *User) (user *User, err error)
	Find(ctx context.Context, id uuid.UUID) (user *User, err error)
	FindBy(ctx context.Context, key, value string) (user *User, err error)
}

type UserUsecase interface {
	Register(ctx context.Context, usr *User, form *http.Request) (res interface{}, err error)
	UpdateUser(ctx context.Context, usr *User, form *http.Request) (res interface{}, err error)
	Login(ctx context.Context, credential *Credential) (res interface{}, err error)
	Logout(ctx context.Context, claims jwt.Claims)
	GetUserById(ctx context.Context, id uuid.UUID) (res interface{}, err error)
}
