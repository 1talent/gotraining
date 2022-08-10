package auth_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/api/handlers/auth"
	"github.com/1talent/gotraining/internal/models"
	"github.com/1talent/gotraining/internal/test"
	"github.com/1talent/gotraining/internal/types"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestPostResgisterSuccess(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		ctx := context.Background()
		username := "usernew@example.com"
		payload := test.GenericPayload{
			"username": username,
			"password": test.PlainTestUserPassword,
		}

		res := test.PerformRequest(t, s, "POST", "/api/v1/auth/register", payload, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.PostLoginResponse
		test.ParseResponseAndValidate(t, res, &response)
		assert.NotEmpty(t, response.AccessToken)
		assert.NotEmpty(t, response.RefreshToken)
		assert.Equal(t, int64(s.Config.Auth.AccessTokenValidity.Seconds()), *response.ExpiresIn)
		assert.Equal(t, auth.TokenTypeBearer, *response.TokenType)

		user, err := models.Users(
			models.UserWhere.Username.EQ(null.StringFrom(username)),
			qm.Load(models.UserRels.AppUserProfile),
			qm.Load(models.UserRels.AccessTokens),
			qm.Load(models.UserRels.RefreshTokens),
		).One(ctx, s.DB)
		assert.NoError(t, err)
		assert.Equal(t, null.StringFrom(username), user.Username)
		assert.Equal(t, true, user.LastAuthenticatedAt.Valid)
		assert.WithinDuration(t, time.Now(), user.LastAuthenticatedAt.Time, time.Second*10)
		assert.EqualValues(t, s.Config.Auth.DefaultUserScopes, user.Scopes)

		assert.NotNil(t, user.R.AppUserProfile)
		assert.Equal(t, false, user.R.AppUserProfile.LegalAcceptedAt.Valid)

		assert.Len(t, user.R.AccessTokens, 1)
		assert.Equal(t, strfmt.UUID4(user.R.AccessTokens[0].Token), *response.AccessToken)
		assert.Len(t, user.R.RefreshTokens, 1)
		assert.Equal(t, strfmt.UUID4(user.R.RefreshTokens[0].Token), *response.RefreshToken)
	})

}
