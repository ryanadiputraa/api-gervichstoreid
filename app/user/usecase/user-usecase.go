package usecase

import (
	"context"
	"net/http"

	logging "github.com/sirupsen/logrus"

	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository domain.IUserRepository
	sessionRead    *dbr.Session
	sessionWrite   *dbr.Session
}

func NewUserUsecase(userRepository domain.IUserRepository, read, write *dbr.Session) domain.IUserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
		sessionRead:    read,
		sessionWrite:   write,
	}
}

func (u *UserUsecase) CreateUser(ctx context.Context, payload domain.UserDTO) (err error) {
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

	payload.Password = string(encryptedPassword)
	err = u.userRepository.CreateUser(ctx, tx, payload)
	if err != nil {
		logging.Error("failed to create user: ", err.Error())
		return err
	}
	tx.Commit()

	return
}

func (u *UserUsecase) GetUserData(ctx context.Context, userID string) (user domain.UserData, err error) {
	userS, err := u.userRepository.GetUserData(ctx, u.sessionRead, userID)
	if err != nil {
		logging.Error("failed to get user data: ", err.Error())
		return
	}

	if userS == nil {
		return user, &wrapper.GenericError{
			HTTPCode: http.StatusNotFound,
			Code:     404,
			Message:  wrapper.BadRequestLabel,
		}
	}

	user = *userS
	return
}
