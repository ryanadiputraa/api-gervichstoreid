package domain

import (
	"context"

	logging "github.com/sirupsen/logrus"
)

const userClaimsCtxKey = "uid"

type userClaimsCtxKeyType string

func WithUserClaims(ctx context.Context, userClaims *Claims) context.Context {
	return context.WithValue(ctx, userClaimsCtxKey, userClaims)
}

func GetUserClaims(ctx context.Context) *Claims {
	userClaims, ok := ctx.Value(userClaimsCtxKey).(*Claims)
	if !ok {
		logging.Error("Fail to decode user claims")
		return nil
	}

	if userClaims == nil {
		logging.Error("User claims is null")
	}
	return userClaims
}
