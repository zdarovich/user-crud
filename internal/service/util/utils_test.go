package util_test

import (
	"binance-order-matcher/internal/model"
	"binance-order-matcher/internal/service/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	ErrInternal = errors.New("internal error")
)

type test struct {
	name   string
	user   *model.User
	result error
}

func TestOrderMatcher_GetMatchedOrder(t *testing.T) {
	t.Parallel()

	tests := []test{
		{
			name: "Validate should return error if 'country' field is empty",
			user: &model.User{
				FirstName: "a",
				LastName:  "a",
				Nickname:  "a",
				Password:  "a",
				Email:     "a",
				Country:   "",
			},
			result: ErrInternal,
		},
		{
			name: "Validate should return error if 'country' field is 1 letter long",
			user: &model.User{
				FirstName: "a",
				LastName:  "a",
				Nickname:  "a",
				Password:  "a",
				Email:     "a",
				Country:   "A",
			},
			result: ErrInternal,
		},
		{
			name: "Validate should return error if 'first_name' field is empty",
			user: &model.User{
				FirstName: "",
				LastName:  "a",
				Nickname:  "a",
				Password:  "a",
				Email:     "a",
				Country:   "EE",
			},
			result: ErrInternal,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			err := util.Validate(tc.user)
			require.Error(t, err, tc.result)
		})
	}
}
