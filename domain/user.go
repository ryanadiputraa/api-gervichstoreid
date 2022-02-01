package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gocraft/dbr/v2"
)

type IAuthUsecase interface {
	Login(ctx context.Context, payload LoginPayload) (*GenerateTokenResponse, error)
	Refresh(ctx context.Context, refreshToken string) (*GenerateTokenResponse, error)
	ChangePassword(ctx context.Context, payload ChangePasswordPayload) error
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, payload UserDTO) error
	GetUserData(ctx context.Context, userID string) (UserData, error)
}

type IUserRepository interface {
	CreateUser(ctx context.Context, tx *dbr.Tx, payload UserDTO) error
	GetUser(ctx context.Context, readSession *dbr.Session, userID string) (*User, error)
	GetUserData(ctx context.Context, readSession *dbr.Session, userID string) (*UserData, error)
}

type User struct {
	ID       string `json:"id" db:"id"`
	Fullname string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type UserDTO struct {
	Fullname string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type UserData struct {
	Fullname string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type GenerateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiredAt    int64  `json:"expired_at"`
	RefreshToken string `json:"refresh_token"`
}
