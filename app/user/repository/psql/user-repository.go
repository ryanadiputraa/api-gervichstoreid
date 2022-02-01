package psql

import (
	"context"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
)

type UserRepository struct {
	sessionRead  *dbr.Session
	sessionWrite *dbr.Session
}

func NewUserRepository(read, write *dbr.Session) domain.IUserRepository {
	return &UserRepository{
		sessionRead:  read,
		sessionWrite: write,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, tx *dbr.Tx, payload domain.UserDTO) (err error) {
	tagList := wrapper.GetStructTagList(domain.User{}, "db")
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		_, err := tx.InsertInto("users").Columns(tagList...).Record(payload).ExecContext(ctx)
		return err
	}, nil)

	if errHystrix != nil {
		return &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    errHystrix.Error(),
		}
	}

	return
}

func (r *UserRepository) GetUser(ctx context.Context, readSession *dbr.Session, userID string) (user *domain.User, err error) {
	tagList := wrapper.GetStructTagList(domain.User{}, "db")
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		_, err := readSession.Select(tagList...).From("users").Where(dbr.Eq("id", userID)).LoadContext(ctx, &user)
		return err
	}, nil)

	if errHystrix != nil {
		return user, &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    errHystrix.Error(),
		}
	}

	return
}

func (r *UserRepository) GetUserData(ctx context.Context, readSession *dbr.Session, userID string) (user *domain.UserData, err error) {
	tagList := wrapper.GetStructTagList(domain.UserData{}, "db")
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		_, err := readSession.Select(tagList...).From("users").Where(dbr.Eq("id", userID)).LoadContext(ctx, &user)
		return err
	}, nil)

	if errHystrix != nil {
		return user, &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    errHystrix.Error(),
		}
	}

	return
}
