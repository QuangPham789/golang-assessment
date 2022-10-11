package relation

import (
	"context"
	"github.com/quangpham789/golang-assessment/utils/db"
	"github.com/stretchr/testify/require"
	"testing"
)

var dbURL = "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable"

func TestRepository_CreateFriendship(t *testing.T) {
	tcs := map[string]struct {
		input1    int
		input2    int
		expResult bool
		expErr    error
	}{
		"success": {
			input1:    4,
			input2:    5,
			expResult: true,
		},
		// TODO: "error duplicate email"

		// TODO: "error duplicate primary_key"
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbURL)
			require.NoError(t, err)
			defer dbConn.Close()
			//defer dbConn.Exec("DELETE FROM public.users;")

			// TODO: Load DB user test sql

			friendshipRepo := NewRelationsRepository(dbConn)
			res, err := friendshipRepo.CreateRelation(ctx, tc.input1, tc.input2)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				tc.expResult = res
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
