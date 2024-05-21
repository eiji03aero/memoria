package handler

import (
	"net/http"
	"testing"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/infra/handler/res"
	"memoria-api/infra/tbl"
	"memoria-api/testutil"
	"memoria-api/util"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

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
				"album_id":   util.StrToNilIfEmpty(test.albumID),
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
