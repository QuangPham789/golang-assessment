package service

import (
	"context"
	"errors"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"testing"
)

func TestService_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input     CreateUserInput
		expResult UserResponse
		expErr    error
	}{
		"success": {
			input: CreateUserInput{
				FirstName: "Quang",
				LastName:  "Pham",
				Email:     "nhutquang23@gmail.com",
				Phone:     "02312545678",
				IsActive:  true,
			},
			expResult: UserResponse{
				ID:        15,
				FirstName: "Quang",
				LastName:  "Pham",
				Email:     "nhutquang23@gmail.com",
				Phone:     "0343450044",
				IsActive:  true,
			},
		},
		"error duplicate email": {
			input: CreateUserInput{
				FirstName: "Quang",
				LastName:  "Pham",
				Email:     "dcthang@gmail.com",
				Phone:     "02312545678",
				IsActive:  true,
			},
			expErr: errors.New("models: unable to insert into users: " +
				"pq: duplicate key value violates unique constraint \"users_pkey\""),
		},
	}

	tcMockUserRepo := map[string]struct {
		result models.User
		err    error
	}{
		"success": {
			result: models.User{
				ID:        15,
				Firstname: "Quang",
				Lastname:  "Pham",
				Email:     "dcthang@gmail.com",
				Phone:     null.StringFrom("0343450044"),
				IsActive:  null.BoolFrom(true),
			},
		},
		"error duplicate email": {
			err: errors.New("models: unable to insert into users: " +
				"pq: duplicate key value violates unique constraint \"users_pkey\""),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockUserRepo)
			mockRepo.On("CreateUser", mock.Anything, mock.Anything).
				Return(tcMockUserRepo[desc].result, tcMockUserRepo[desc].err)

			userService := UserService{mockRepo}
			res, err := userService.CreateUser(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}
