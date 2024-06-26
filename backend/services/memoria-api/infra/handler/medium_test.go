package handler

import (
	"net/http"
	"testing"

	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/domain/value"
	"memoria-api/infra/handler/res"
	"memoria-api/infra/tbl"
	"memoria-api/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMediumFind_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		query    string
		preAmrs  func(e testutil.UserEnv) []*tbl.AlbumMediumRelation
		preMedia func(e testutil.UserEnv) []*tbl.Medium
		expected func(e testutil.UserEnv) MediumFindRes
	}{
		{
			name:  "Case 1: Make sure only user space",
			query: "",
			preAmrs: func(e testutil.UserEnv) []*tbl.AlbumMediumRelation {
				return []*tbl.AlbumMediumRelation{}
			},
			preMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{ID: "m1", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "file 1", Extension: ".png"},
					{ID: "m2", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "file 2", Extension: ".png"},
					{ID: "m3", UserID: e.User.ID, UserSpaceID: "other", Name: "file 3", Extension: ".png"},
				}
			},
			expected: func(e testutil.UserEnv) MediumFindRes {
				return MediumFindRes{
					Media: []*res.Medium{
						{ID: "m1", Name: "file 1", Extension: ".png"},
						{ID: "m2", Name: "file 2", Extension: ".png"},
					},
				}
			},
		},
		{
			name:  "Case 2: Filter by album id",
			query: "album_id=a1",
			preAmrs: func(e testutil.UserEnv) []*tbl.AlbumMediumRelation {
				return []*tbl.AlbumMediumRelation{
					{AlbumID: "a1", MediumID: "m1"},
				}
			},
			preMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{ID: "m1", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "file 1", Extension: ".png"},
					{ID: "m2", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "file 2", Extension: ".png"},
				}
			},
			expected: func(e testutil.UserEnv) MediumFindRes {
				return MediumFindRes{
					Media: []*res.Medium{
						{ID: "m1", Name: "file 1", Extension: ".png"},
					},
				}
			},
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		t.Log(test.name)
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContext(
			http.MethodPost,
			"/?"+test.query,
		)
		env.SetupAuthorization(c)

		api.DB().Create(test.preAmrs(env))
		api.DB().Create(test.preMedia(env))

		// -------------------- execution --------------------
		mediumH := NewMedium()
		status, data, err := mediumH.Find(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := MediumFindRes{}
		verifyJSONEncoding(data, &decodedRes)

		expected := test.expected(env)
		assert.Equal(t, len(expected.Media), len(decodedRes.Media))
		for i := range expected.Media {
			assert.Equal(t, expected.Media[i].ID, decodedRes.Media[i].ID)
			assert.Equal(t, expected.Media[i].Name, decodedRes.Media[i].Name)
			assert.Equal(t, expected.Media[i].Extension, decodedRes.Media[i].Extension)
			assert.NotEmpty(t, decodedRes.Media[i].OriginalURL)
		}
	}
}

func TestMediumFindOne_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		expRes func(e testutil.UserEnv) MediumFindOneRes
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Finds one",
				Seeder: func(e testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{
							{ID: "m1", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
						},
					}
				},
			},
			expRes: func(e testutil.UserEnv) MediumFindOneRes {
				return MediumFindOneRes{
					Medium: &res.Medium{
						ID: "m1", Name: "f1", Extension: ".png",
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
		c.Params = gin.Params{{Key: "id", Value: "m1"}}

		test.InstallSeeds(api.DB(), env)

		// -------------------- execution --------------------
		mediumH := NewMedium()
		status, data, err := mediumH.FindOne(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := MediumFindOneRes{}
		verifyJSONEncoding(data, &decodedRes)

		expRes := test.expRes(env)
		assert.Equal(t, expRes.Medium.ID, decodedRes.Medium.ID)
		assert.Equal(t, expRes.Medium.Name, decodedRes.Medium.Name)
		assert.Equal(t, expRes.Medium.Extension, decodedRes.Medium.Extension)
	}
}

func TestMediumGetPage_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		query  string
		expRes func(e testutil.UserEnv) MediumGetPageRes
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Gets page of all",
				Seeder: func(e testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{
							{ID: "m1", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
							{ID: "m2", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
							{ID: "m3", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
						},
					}
				},
			},
			query: "medium_id=m2",
			expRes: func(e testutil.UserEnv) MediumGetPageRes {
				return MediumGetPageRes{
					Pagi: res.Pagination{
						CurrentPage: 2,
						PerPage:     1,
						TotalPage:   3,
					},
				}
			},
		},
		{
			TestCase: testutil.TestCase{
				Name: "Gets page of album",
				Seeder: func(e testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{
							{ID: "m1", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
							{ID: "m2", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
							{ID: "m3", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
						},
						[]*tbl.Album{
							{ID: "a1", Name: "a1"},
						},
						[]*tbl.AlbumMediumRelation{
							{AlbumID: "a1", MediumID: "m2"},
							{AlbumID: "a1", MediumID: "m3"},
						},
					}
				},
			},
			query: "medium_id=m2&album_id=a1",
			expRes: func(e testutil.UserEnv) MediumGetPageRes {
				return MediumGetPageRes{
					Pagi: res.Pagination{
						CurrentPage: 1,
						PerPage:     1,
						TotalPage:   2,
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
			"/?"+test.query,
		)
		env.SetupAuthorization(c)

		test.InstallSeeds(api.DB(), env)

		// -------------------- execution --------------------
		mediumH := NewMedium()
		status, data, err := mediumH.GetPage(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := MediumGetPageRes{}
		verifyJSONEncoding(data, &decodedRes)

		expRes := test.expRes(env)
		assert.Equal(t, expRes.Pagi.CurrentPage, decodedRes.Pagi.CurrentPage)
		assert.Equal(t, expRes.Pagi.PerPage, decodedRes.Pagi.PerPage)
		assert.Equal(t, expRes.Pagi.TotalPage, decodedRes.Pagi.TotalPage)
	}
}

func TestMediumDelete_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		id       string
		expMedia func(e testutil.UserEnv) []*tbl.Medium
		expAmrs  func(e testutil.UserEnv) []*tbl.AlbumMediumRelation
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Deletes medium",
				Seeder: func(env testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{
							{ID: "m1"},
						},
					}
				},
			},
			id: "m1",
			expMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{}
			},
		},
		{
			TestCase: testutil.TestCase{
				Name: "Deletes medium",
				Seeder: func(env testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{
							{ID: "m1"},
							{ID: "m2"},
						},
						[]*tbl.Album{
							{ID: "a1"},
						},
						[]*tbl.AlbumMediumRelation{
							{AlbumID: "a1", MediumID: "m1"},
						},
					}
				},
			},
			id: "m1",
			expMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{ID: "m2"},
				}
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
		test.InstallSeeds(api.DB(), env)

		c := newGinContext(
			http.MethodPost,
			"/",
		)
		c.Params = gin.Params{{Key: "id", Value: test.id}}
		env.SetupAuthorization(c)

		// -------------------- execution --------------------
		mediumH := NewMedium()
		_, _, err := mediumH.Delete(c, reg)

		// -------------------- assertion --------------------
		assert.NoError(t, err)

		actual, err := reg.NewMediumRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)

		expected := test.expMedia(env)
		assert.Equal(t, len(expected), len(actual))

		if test.expAmrs != nil {
			actual, err := reg.NewAlbumMediumRelationRepository().Find(&repository.FindOption{})
			assert.NoError(t, err)

			expected := test.expAmrs(env)
			assert.Equal(t, len(expected), len(actual))
		}
	}
}

func TestMediumRequestUploadURLs_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		fileNames     []string
		albumID       string
		preAlbums     func(e testutil.UserEnv) []*tbl.Album
		expectedMedia func(e testutil.UserEnv) []*tbl.Medium
	}{
		{
			name:      "Case 1: Create multiple",
			fileNames: []string{"Hoge.png", "Family.jpeg"},
			albumID:   "",
			expectedMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "Hoge", Extension: ".png"},
					{UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "Family", Extension: ".jpeg"},
				}
			},
		},
		{
			name:      "Case 2: Create and link to album",
			fileNames: []string{"Hoge.png"},
			albumID:   "album-id",
			preAlbums: func(e testutil.UserEnv) []*tbl.Album {
				return []*tbl.Album{
					{ID: "album-id", Name: "1"},
				}
			},
			expectedMedia: func(e testutil.UserEnv) []*tbl.Medium {
				return []*tbl.Medium{
					{UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "Hoge", Extension: ".png"},
				}
			},
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		t.Log(test.name)
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"file_names": test.fileNames,
				"album_ids":  []string{test.albumID},
			},
		)
		env.SetupAuthorization(c)

		if test.preAlbums != nil {
			err = api.DB().Create(test.preAlbums(env)).Error
			assert.NoError(t, err)
		}

		api.MockS3Client.EXPECT().GetPresignedPutObjectURL(gomock.Any()).Return("https://hoge.com", nil).AnyTimes()

		// -------------------- execution --------------------
		mediumH := NewMedium()
		status, data, err := mediumH.RequestUploadURLs(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := MediumRequestUploadURLsRes{}
		verifyJSONEncoding(data, &decodedRes)

		// check res
		for _, uploadURL := range decodedRes.UploadURLs {
			assert.NotEmpty(t, uploadURL.URL)
		}
		// medium record created
		media, err := reg.NewMediumRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, len(test.fileNames), len(media))
		for i, medium := range test.expectedMedia(env) {
			assert.Equal(t, env.User.ID, media[i].UserID)
			assert.Equal(t, env.UserSpace.ID, media[i].UserSpaceID)
			assert.Equal(t, medium.Name, media[i].Name)
			assert.Equal(t, medium.Extension, media[i].Extension)
		}
		// check if linked to album
		if test.albumID != "" {
			amrs, err := reg.NewAlbumMediumRelationRepository().Find(&repository.FindOption{
				Filters: []*repository.FindOptionFilter{
					{Query: "album_id = ?", Value: test.albumID},
				},
			})
			assert.NoError(t, err)
			assert.Equal(t, len(test.fileNames), len(amrs))
		}
	}
}

func TestMediumConfirmUploads_S_Invoker(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		body    map[string]any
		expUsas func(e testutil.UserEnv) []model.UserSpaceActivity
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Gets page of all",
				Seeder: func(e testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{
							{ID: "m1", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
							{ID: "m2", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
							{ID: "m3", UserID: e.User.ID, UserSpaceID: e.UserSpace.ID, Name: "f1", Extension: ".png"},
						},
					}
				},
			},
			body: map[string]any{
				"medium_ids": []string{"m1", "m2", "m3"},
			},
			expUsas: func(e testutil.UserEnv) []model.UserSpaceActivity {
				return []model.UserSpaceActivity{
					{UserSpaceID: e.UserSpace.ID, Type: value.UserSpaceActivityType_UserUploadedMedia, Data: `{"user_id": "` + e.User.ID + `", "medium_ids": ["m1", "m2", "m3"]}`},
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

		test.InstallSeeds(api.DB(), env)

		ids := test.body["medium_ids"].([]string)
		for _, id := range ids {
			api.MockBGJobInvoker.
				EXPECT().
				CreateThumbnails(gomock.Cond(func(dto any) bool {
					return (dto).(interfaces.BGJobInvokerCreateThumbnailsDTO).MediumID == id
				}))
		}

		// -------------------- execution --------------------
		mediumH := NewMedium()
		_, _, err := mediumH.ConfirmUploads(c, reg)

		// -------------------- assertion --------------------
		assert.NoError(t, err)

		expUsas := test.expUsas(env)
		actual, err := reg.NewUserSpaceActivityRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)

		assert.Equal(t, len(expUsas), len(actual))
		for i := range expUsas {
			assert.Equal(t, expUsas[i].Type, actual[i].Type)
			assert.Equal(t, expUsas[i].UserSpaceID, actual[i].UserSpaceID)
			assert.Equal(t, expUsas[i].Data, actual[i].Data)
		}
	}
}
