package handler

import (
	"net/http"
	"testing"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/infra/handler/res"
	"memoria-api/infra/tbl"
	"memoria-api/testutil"
	"memoria-api/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAlbumCreate_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name string
	}{
		{
			name: "Baby memory",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"name": util.StrToNilIfEmpty(test.name),
			},
		)
		env.SetupAuthorization(c)

		// -------------------- execution --------------------
		albumH := NewAlbum()
		status, data, err := albumH.Create(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := AlbumCreateRes{}
		verifyJSONEncoding(data, &decodedRes)

		// album record created
		albums, err := reg.NewAlbumRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(albums))
		album := albums[0]
		assert.Equal(t, test.name, album.Name)
		// user space album relation created
		usars, err := reg.NewUserSpaceAlbumRelationRepository().Find(&repository.FindOption{})
		assert.Equal(t, 1, len(usars))
		usar := usars[0]
		assert.Equal(t, env.UserSpace.ID, usar.UserSpaceID)
		assert.Equal(t, decodedRes.Album.ID, usar.AlbumID)
	}
}

func TestAlbumCreate_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name           string
		validationKey  cerrors.ValidationKey
		validationName string
	}{
		{
			name:           "",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "name",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"name": util.StrToNilIfEmpty(test.name),
			},
		)
		env.SetupAuthorization(c)

		// -------------------- execution --------------------
		albumH := NewAlbum()
		_, _, err := albumH.Create(c, reg)

		// -------------------- assertion --------------------
		validationErr, ok := err.(cerrors.Validation)
		assert.True(t, ok)

		assert.Equal(t, test.validationKey, validationErr.Key)
		assert.Equal(t, test.validationName, validationErr.Name)
	}
}

func TestAlbumFind_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		expected  func(e testutil.UserEnv) AlbumFindRes
	}{
		{
			preAlbums: func(e testutil.UserEnv) []*tbl.Album {
				return []*tbl.Album{
					{ID: "1", Name: "a1"},
					{ID: "2", Name: "a2"},
					{ID: "3", Name: "a3"},
				}
			},
			preUsars: func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation {
				return []*tbl.UserSpaceAlbumRelation{
					{UserSpaceID: e.UserSpace.ID, AlbumID: "1"},
					{UserSpaceID: e.UserSpace.ID, AlbumID: "2"},
					{UserSpaceID: "other", AlbumID: "3"},
				}
			},
			expected: func(e testutil.UserEnv) AlbumFindRes {
				return AlbumFindRes{
					Albums: []*res.Album{
						{ID: "1", Name: "a1"},
						{ID: "2", Name: "a2"},
					},
				}
			},
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContext(
			http.MethodPost,
			"/",
		)
		env.SetupAuthorization(c)

		api.DB().Create(test.preAlbums(env))
		api.DB().Create(test.preUsars(env))

		// -------------------- execution --------------------
		albumH := NewAlbum()
		status, data, err := albumH.Find(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := AlbumFindRes{}
		verifyJSONEncoding(data, &decodedRes)

		expected := test.expected(env)
		assert.Equal(t, len(expected.Albums), len(decodedRes.Albums))
		for i := range expected.Albums {
			assert.Equal(t, expected.Albums[i].ID, decodedRes.Albums[i].ID)
			assert.Equal(t, expected.Albums[i].Name, decodedRes.Albums[i].Name)
		}
	}
}

func TestAlbumFindOne_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		expected  func(e testutil.UserEnv) AlbumFindOneRes
	}{
		{
			preAlbums: func(e testutil.UserEnv) []*tbl.Album {
				return []*tbl.Album{
					{ID: "1", Name: "a1"},
				}
			},
			preUsars: func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation {
				return []*tbl.UserSpaceAlbumRelation{
					{UserSpaceID: e.UserSpace.ID, AlbumID: "1"},
					{UserSpaceID: e.UserSpace.ID, AlbumID: "2"},
				}
			},
			expected: func(e testutil.UserEnv) AlbumFindOneRes {
				return AlbumFindOneRes{
					Album: res.Album{
						ID: "1", Name: "a1",
					},
				}
			},
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContext(
			http.MethodGet,
			"/api/auth/albums/1",
		)
		c.Params = []gin.Param{{Key: "id", Value: "1"}}
		env.SetupAuthorization(c)

		api.DB().Create(test.preAlbums(env))
		api.DB().Create(test.preUsars(env))

		// -------------------- execution --------------------
		albumH := NewAlbum()
		status, data, err := albumH.FindOne(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := AlbumFindOneRes{}
		verifyJSONEncoding(data, &decodedRes)

		expected := test.expected(env)
		assert.Equal(t, expected.Album.ID, decodedRes.Album.ID)
		assert.Equal(t, expected.Album.Name, decodedRes.Album.Name)
	}
}
