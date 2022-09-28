package util

import (
	"binance-order-matcher/internal/model"
	"github.com/pkg/errors"
)

func Validate(user *model.User) error {
	if user == nil {
		return nil
	}
	if len(user.Country) != 2 {
		return errors.New("'country' field must be 2 letters long")
	} else if len(user.Email) == 0 {
		return errors.New("'email' field must not be empty")
	} else if len(user.Password) == 0 {
		return errors.New("'password' field must not be empty")
	} else if len(user.Nickname) == 0 {
		return errors.New("'nickname' field must not be empty")
	} else if len(user.FirstName) == 0 {
		return errors.New("'first_name' field must not be empty")
	} else if len(user.LastName) == 0 {
		return errors.New("'last_name' field must not be empty")
	}
	return nil
}
