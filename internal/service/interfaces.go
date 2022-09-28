//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=service_test
package service

import (
	"binance-order-matcher/internal/model"
	"context"
)

type (
	UserRepo interface {
		Save(ctx context.Context, user *model.User) error
		Update(ctx context.Context, user *model.User) error
		Delete(ctx context.Context, user *model.User) error
		Get(ctx context.Context, page, limit int, filter model.User) ([]*model.User, error)
	}
)
