package usecase

import (
	"context"
	"net/http"

	logging "github.com/sirupsen/logrus"

	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
)

type UserUsecase struct {
	sessionRead    *dbr.Session
	sessionWrite   *dbr.Session
	userRepository domain.IUserRepository
}

func NewUserUsecase(read, write *dbr.Session, userRepository domain.IUserRepository) domain.IUserUsecase {
	return &UserUsecase{
		sessionRead:    read,
		sessionWrite:   write,
		userRepository: userRepository,
	}
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
