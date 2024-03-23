package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_database "github.com/arya2004/Xyfin/database/mock"
	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAccount(t *testing.T) {
	testCases := []struct {
		name         string
		accountID    int64
		buildStubs   func(store *mock_database.MockStore)
		expectStatus int
	}{
		{
			name:      "OK",
			accountID: 1,
			buildStubs: func(store *mock_database.MockStore) {
				const id int64 = 1
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(id)).
					Return(randomAccount(id), nil).
					Times(1)
			},
			expectStatus: http.StatusOK,
		},
		{
			name:      "NotFound",
			accountID: 2,
			buildStubs: func(store *mock_database.MockStore) {
				const id int64 = 2
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(id)).
					Return(database.Account{}, sql.ErrNoRows).
					Times(1)
			},
			expectStatus: http.StatusNotFound,
		},
		{
			name:      "InternalError",
			accountID: 3,
			buildStubs: func(store *mock_database.MockStore) {
				const id int64 = 3
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(id)).
					Return(database.Account{}, sql.ErrConnDone).
					Times(1)
			},
			expectStatus: http.StatusInternalServerError,
		},
		{
			name:      "BadRequest",
			accountID: 0,
			buildStubs: func(store *mock_database.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_database.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}

func randomAccount(id int64) database.Account {
	return database.Account{
		ID:        id,
		Owner:     util.RandomOwner(),
		Balance:   util.RandomMoney(),
		Currency:  util.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}