package api

import (
	"os"
	"testing"
	"time"

	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store database.Store) *Server {
	config := util.Configuration{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}