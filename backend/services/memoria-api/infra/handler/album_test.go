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
		testutil.TestCase
		name string
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Creates album",
			},
			name: "Baby memory",
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)
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
		testutil.TestCase
		name           string
		validationKey  cerrors.ValidationKey
		validationName string
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Validates existence of name",
			},
			name:           "",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "name",
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)

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
		test.AssertValidationError(testutil.TestCaseAssertValidationErrorDTO{
			T:        t,
			ExpVKey:  test.validationKey,
			ExpVName: test.validationName,
			Err:      err,
		})
	}
}

func TestAlbumFind_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		expected  func(e testutil.UserEnv) AlbumFindRes
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Finds",
			},
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

	for i, test := range tests {
		test.LogCase(t, i)

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
		testutil.TestCase
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		expected  func(e testutil.UserEnv) AlbumFindOneRes
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Finds one",
			},
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
					Album: &res.Album{
						ID: "1", Name: "a1",
					},
				}
			},
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)

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

func TestAlbumAddMedia_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		preMedia  func(e testutil.UserEnv) []*tbl.Medium
		body      map[string]any
		expected  func(e testutil.UserEnv) []*tbl.AlbumMediumRelation
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Add Multiple media",
			},
			preAlbums: func(e testutil.UserEnv) []*tbl.Album {
				return []*tbl.Album{
					{ID: "a1", Name: "a1"},
				}
			},
			preUsars: func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation {
				return []*tbl.UserSpaceAlbumRelation{
					{UserSpaceID: e.UserSpace.ID, AlbumID: "a1"},
				}
			},
			preMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{ID: "m1", UserSpaceID: e.UserSpace.ID},
					{ID: "m2", UserSpaceID: e.UserSpace.ID},
				}
			},
			body: map[string]any{
				"album_ids":  []string{"a1"},
				"medium_ids": []string{"m1", "m2"},
			},
			expected: func(e testutil.UserEnv) []*tbl.AlbumMediumRelation {
				return []*tbl.AlbumMediumRelation{
					{AlbumID: "a1", MediumID: "m1"},
					{AlbumID: "a1", MediumID: "m2"},
				}
			},
		},
		{
			TestCase: testutil.TestCase{
				Name: "Add multiple media to multiple albums",
			},
			preAlbums: func(e testutil.UserEnv) []*tbl.Album {
				return []*tbl.Album{
					{ID: "a1", Name: "a1"},
					{ID: "a2", Name: "a2"},
				}
			},
			preUsars: func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation {
				return []*tbl.UserSpaceAlbumRelation{
					{UserSpaceID: e.UserSpace.ID, AlbumID: "a1"},
					{UserSpaceID: e.UserSpace.ID, AlbumID: "a2"},
				}
			},
			preMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{ID: "m1", UserSpaceID: e.UserSpace.ID},
					{ID: "m2", UserSpaceID: e.UserSpace.ID},
				}
			},
			body: map[string]any{
				"album_ids":  []string{"a1", "a2"},
				"medium_ids": []string{"m1", "m2"},
			},
			expected: func(e testutil.UserEnv) []*tbl.AlbumMediumRelation {
				return []*tbl.AlbumMediumRelation{
					{AlbumID: "a1", MediumID: "m1"},
					{AlbumID: "a1", MediumID: "m2"},
					{AlbumID: "a2", MediumID: "m1"},
					{AlbumID: "a2", MediumID: "m2"},
				}
			},
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)

		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			test.body,
		)
		env.SetupAuthorization(c)

		api.DB().Create(test.preAlbums(env))
		api.DB().Create(test.preUsars(env))
		api.DB().Create(test.preMedia(env))

		// -------------------- execution --------------------
		albumH := NewAlbum()
		status, _, err := albumH.AddMedia(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		actual, err := reg.NewAlbumMediumRelationRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)

		expected := test.expected(env)
		assert.Equal(t, len(expected), len(actual))
		for i := range expected {
			assert.Equal(t, expected[i].AlbumID, actual[i].AlbumID)
			assert.Equal(t, expected[i].MediumID, actual[i].MediumID)
		}
	}
}

func TestAlbumAddMedia_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		preMedia  func(e testutil.UserEnv) []*tbl.Medium
		body      map[string]any
		expVKey   cerrors.ValidationKey
		expVName  string
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Not to add media album that belong to other user space",
			},
			preAlbums: func(e testutil.UserEnv) []*tbl.Album {
				return []*tbl.Album{
					{ID: "a1", Name: "a1"},
				}
			},
			preUsars: func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation {
				return []*tbl.UserSpaceAlbumRelation{
					{UserSpaceID: "other", AlbumID: "a1"},
				}
			},
			preMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{ID: "m1", UserSpaceID: "other"},
				}
			},
			body: map[string]any{
				"album_ids":  []string{"a1"},
				"medium_ids": []string{"m1"},
			},
			expVKey:  cerrors.ValidationKey_Consistency,
			expVName: "user-space-id",
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)

		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			test.body,
		)
		env.SetupAuthorization(c)

		api.DB().Create(test.preAlbums(env))
		api.DB().Create(test.preUsars(env))
		api.DB().Create(test.preMedia(env))

		// -------------------- execution --------------------
		albumH := NewAlbum()
		_, _, err := albumH.AddMedia(c, reg)

		// -------------------- assertion --------------------
		test.AssertValidationError(testutil.TestCaseAssertValidationErrorDTO{
			T:        t,
			ExpVKey:  test.expVKey,
			ExpVName: test.expVName,
			Err:      err,
		})
	}
}

func TestAlbumRemoveMedia_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		body    map[string]any
		expAmrs func(e testutil.UserEnv) []*tbl.AlbumMediumRelation
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Remove Multiple media",
				Seeder: func(e testutil.UserEnv) []any {
					return []any{
						[]*tbl.Album{
							{ID: "a1", Name: "a1"},
						},
						[]*tbl.UserSpaceAlbumRelation{
							{UserSpaceID: e.UserSpace.ID, AlbumID: "a1"},
						},
						[]*tbl.Medium{
							{ID: "m1", UserSpaceID: e.UserSpace.ID},
							{ID: "m2", UserSpaceID: e.UserSpace.ID},
							{ID: "m3", UserSpaceID: e.UserSpace.ID},
						},
						[]*tbl.AlbumMediumRelation{
							{AlbumID: "a1", MediumID: "m1"},
							{AlbumID: "a1", MediumID: "m2"},
							{AlbumID: "a1", MediumID: "m3"},
						},
					}
				},
			},
			body: map[string]any{
				"album_ids":  []string{"a1"},
				"medium_ids": []string{"m1", "m2"},
			},
			expAmrs: func(e testutil.UserEnv) []*tbl.AlbumMediumRelation {
				return []*tbl.AlbumMediumRelation{
					{AlbumID: "a1", MediumID: "m3"},
				}
			},
		},
		{
			TestCase: testutil.TestCase{
				Name: "Remove multiple media from multiple albums",
				Seeder: func(e testutil.UserEnv) []any {
					return []any{
						[]*tbl.Album{
							{ID: "a1", Name: "a1"},
							{ID: "a2", Name: "a2"},
						},
						[]*tbl.UserSpaceAlbumRelation{
							{UserSpaceID: e.UserSpace.ID, AlbumID: "a1"},
							{UserSpaceID: e.UserSpace.ID, AlbumID: "a2"},
						},
						[]*tbl.Medium{
							{ID: "m1", UserSpaceID: e.UserSpace.ID},
							{ID: "m2", UserSpaceID: e.UserSpace.ID},
						},
						[]*tbl.AlbumMediumRelation{
							{AlbumID: "a1", MediumID: "m1"},
							{AlbumID: "a2", MediumID: "m2"},
						},
					}
				},
			},
			body: map[string]any{
				"album_ids":  []string{"a1", "a2"},
				"medium_ids": []string{"m1", "m2"},
			},
			expAmrs: func(e testutil.UserEnv) []*tbl.AlbumMediumRelation {
				return []*tbl.AlbumMediumRelation{}
			},
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)

		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			test.body,
		)
		env.SetupAuthorization(c)

		test.InstallSeeds(api.DB(), env)

		// -------------------- execution --------------------
		albumH := NewAlbum()
		status, _, err := albumH.RemoveMedia(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		actual, err := reg.NewAlbumMediumRelationRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)

		expected := test.expAmrs(env)
		assert.Equal(t, len(expected), len(actual))
		for i := range expected {
			assert.Equal(t, expected[i].AlbumID, actual[i].AlbumID)
			assert.Equal(t, expected[i].MediumID, actual[i].MediumID)
		}
	}
}

func TestAlbumRemoveMedia_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		preAlbums func(e testutil.UserEnv) []*tbl.Album
		preUsars  func(e testutil.UserEnv) []*tbl.UserSpaceAlbumRelation
		preMedia  func(e testutil.UserEnv) []*tbl.Medium
		body      map[string]any
		expVKey   cerrors.ValidationKey
		expVName  string
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Not to add media album that belong to other user space",
				Seeder: func(env testutil.UserEnv) []any {
					return []any{
						[]*tbl.Album{
							{ID: "a1", Name: "a1"},
						},
						[]*tbl.UserSpaceAlbumRelation{
							{UserSpaceID: "other", AlbumID: "a1"},
						},
						[]*tbl.Medium{
							{ID: "m1", UserSpaceID: "other"},
						},
					}
				},
			},
			body: map[string]any{
				"album_ids":  []string{"a1"},
				"medium_ids": []string{"m1"},
			},
			expVKey:  cerrors.ValidationKey_Consistency,
			expVName: "user-space-id",
		},
	}

	for i, test := range tests {
		test.LogCase(t, i)

		// -------------------- preparation --------------------
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			test.body,
		)
		env.SetupAuthorization(c)

		test.InstallSeeds(api.DB(), env)

		// -------------------- execution --------------------
		albumH := NewAlbum()
		_, _, err := albumH.AddMedia(c, reg)

		// -------------------- assertion --------------------
		test.AssertValidationError(testutil.TestCaseAssertValidationErrorDTO{
			T:        t,
			ExpVKey:  test.expVKey,
			ExpVName: test.expVName,
			Err:      err,
		})
	}
}
