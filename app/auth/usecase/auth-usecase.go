package usecase

import (
	"context"
	"net/http"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
	logging "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	sessionRead    *dbr.Session
	sessionWrite   *dbr.Session
	userRepository domain.IUserRepository
}

func NewAuthUsecase(read, write *dbr.Session, userRepository domain.IUserRepository) domain.IAuthUsecase {
	return &AuthUsecase{
		sessionRead:    read,
		sessionWrite:   write,
		userRepository: userRepository,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, payload domain.UserDTO) (err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    err.Error(),
		}
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    err.Error(),
		}
	}

	tx, err := u.sessionWrite.BeginTx(ctx, nil)
	if err != nil {
		logging.Error("failed to begin db transactions: " + err.Error())
		return err
	}
	defer tx.RollbackUnlessCommitted()

	payload.ID = id
	payload.Password = string(encryptedPassword)
	err = u.userRepository.CreateUser(ctx, tx, payload)
	if err != nil {
		logging.Error("failed to create user: ", err.Error())
		return err
	}
	tx.Commit()
	return
}

func (u *AuthUsecase) Login(ctx context.Context, payload domain.LoginDTO) (tokens *domain.TokenResponse, err error) {
	return
}

func (u *AuthUsecase) Refresh(ctx context.Context, refreshToken string) (tokens *domain.TokenResponse, err error) {
	return
}

// func (u *AuthUsecase) ChangePassword(ctx context.Context, payload domain.ChangePasswordPayload) (err error) {
// 	return
// }
