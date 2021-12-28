package domain

import "github.com/dgrijalva/jwt-go"

// type IAuthUsecase struct {
// 	Login()
// }

type User struct {
	ID       string `json:"id" db:"id"`
	Fullname string `json:"fullname" db:"fullname"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type UserDetail struct {
	Fullname string `json:"fullname" db:"fullname"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	UserID          string `json:"user_id"`
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
