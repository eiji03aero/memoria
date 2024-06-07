package handler

import (
	"net/http"
	"testing"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/value"
	"memoria-api/infra/handler/res"
	"memoria-api/infra/tbl"
	"memoria-api/testutil"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTimelineFind_S_Timeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		query        string
		exp          []map[string]any
		expContent   string
		expMediumIDs []string
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Gets timeline units",
				Seeder: func(env testutil.UserEnv) []any {
					return []any{
						[]*tbl.TimelinePost{
							{ID: "tu1", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID},
						},
						[]*tbl.UserSpaceActivity{
							{ID: "tu2", UserSpaceID: env.UserSpace.ID, Type: string(value.UserSpaceActivityType_InviteUserJoined), Data: "{}"},
							{ID: "tu3", UserSpaceID: env.UserSpace.ID, Type: string(value.UserSpaceActivityType_UserUploadedMedia), Data: "{}"},
						},
						[]*tbl.Thread{
							{ID: "t1", UserSpaceID: env.UserSpace.ID},
						},
						[]*tbl.TimelinePostThreadRelation{
							{TimelinePostID: "tu1", ThreadID: "t1"},
						},
						[]*tbl.MicroPost{
							{ID: "mp1", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID, Content: "Hoge"},
						},
						[]*tbl.ThreadMicroPostRelation{
							{ThreadID: "t1", MicroPostID: "mp1"},
						},
						[]*tbl.Medium{
							{ID: "m1", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID, Name: "Hoge", Extension: ".png"},
						},
						[]*tbl.MicroPostMediumRelation{
							{MicroPostID: "mp1", MediumID: "m1"},
						},
					}
				},
			},
			query: "",
			exp: []map[string]any{
				{
					"type": "user-space-activity",
				},
				{
					"type": "user-space-activity",
				},
				{
					"type":      "timeline-post",
					"content":   "Hoge",
					"mediumIDs": []string{"m1"},
				},
			},
		},
	}

	for i, test := range tests {
		// -------------------- preparation --------------------
		test.LogCase(t, i)
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContext(
			http.MethodGet,
			"/?"+test.query,
		)
		env.SetupAuthorization(c)

		test.InstallSeeds(api.DB(), env)

		// -------------------- execution --------------------
		timelineH := NewTimeline()
		status, data, err := timelineH.Find(c, reg)

		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := TimelineFindRes{}
		verifyJSONEncoding(data, &decodedRes)

		decodedRes = data.(TimelineFindRes)
		assert.Equal(t, len(test.exp), len(decodedRes.Units))
		for i, unit := range decodedRes.Units {
			exp := test.exp[i]
			assert.Equal(t, exp["type"], unit.Type)

			if content, ok := exp["content"]; ok {
				actual := unit.Data.(*res.TimelinePost).Thread.MicroPosts[0].Content
				assert.Equal(t, content, actual)
			}
		}
	}
}

func TestTimelineFind_S_Pagination(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	seeder := func(env testutil.UserEnv) []any {
		return []any{
			[]*tbl.TimelinePost{
				{ID: "tp1", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID},
				{ID: "tp2", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID},
				{ID: "tp3", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID},
				{ID: "tp4", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID},
				{ID: "tp5", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID},
			},
		}
	}

	tests := []struct {
		testutil.TestCase
		query         string
		exp           []string
		expNextCursor string
		expPrevCursor string
	}{
		{
			TestCase: testutil.TestCase{
				Name:   "Find with after",
				Seeder: seeder,
			},
			query:         "cursor=tp2&cafter=3",
			exp:           []string{"tp2", "tp1"},
			expNextCursor: "tp1",
			expPrevCursor: "tp2",
		},
		{
			TestCase: testutil.TestCase{
				Name:   "Find after with exclude",
				Seeder: seeder,
			},
			query:         "cursor=tp3&cafter=3&cexclude=true",
			exp:           []string{"tp2", "tp1"},
			expNextCursor: "tp1",
			expPrevCursor: "tp2",
		},
		{
			TestCase: testutil.TestCase{
				Name:   "Find with before",
				Seeder: seeder,
			},
			query:         "cursor=tp1&cbefore=3",
			exp:           []string{"tp3", "tp2", "tp1"},
			expNextCursor: "tp1",
			expPrevCursor: "tp3",
		},
		{
			TestCase: testutil.TestCase{
				Name:   "Find before with exclude",
				Seeder: seeder,
			},
			query:         "cursor=tp3&cbefore=3&cexclude=true",
			exp:           []string{"tp5", "tp4"},
			expNextCursor: "tp4",
			expPrevCursor: "tp5",
		},
		{
			TestCase: testutil.TestCase{
				Name:   "Find no more on initial edge",
				Seeder: seeder,
			},
			query:         "cursor=tp1&cafter=3&cexclude=true",
			exp:           []string{},
			expNextCursor: "tp1",
			expPrevCursor: "tp1",
		},
		{
			TestCase: testutil.TestCase{
				Name:   "Find no more on end edge",
				Seeder: seeder,
			},
			query:         "cursor=tp5&cbefore=3&cexclude=true",
			exp:           []string{},
			expNextCursor: "tp5",
			expPrevCursor: "tp5",
		},
	}

	for i, test := range tests {
		// -------------------- preparation --------------------
		test.LogCase(t, i)
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContext(
			http.MethodGet,
			"/?"+test.query,
		)
		env.SetupAuthorization(c)

		test.InstallSeeds(api.DB(), env)

		// -------------------- execution --------------------
		timelineH := NewTimeline()
		status, data, err := timelineH.Find(c, reg)

		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := TimelineFindRes{}
		verifyJSONEncoding(data, &decodedRes)

		assert.Equal(t, len(test.exp), len(decodedRes.Units))
		assert.Equal(t, test.expNextCursor, decodedRes.CPagi.NextCursor)
		assert.Equal(t, test.expPrevCursor, decodedRes.CPagi.PrevCursor)
		for i, unit := range decodedRes.Units {
			assert.Equal(t, test.exp[i], unit.ID)
		}
	}
}

func TestTimelineCreatePost_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		testutil.TestCase
		content   string
		MediumIDs []string
	}{
		{
			TestCase: testutil.TestCase{
				Name: "Creates timeline post",
			},
			content: "Hoge dayo",
		},
		{
			TestCase: testutil.TestCase{
				Name: "Creates timeline post with media",
				Seeder: func(env testutil.UserEnv) []any {
					return []any{
						[]*tbl.Medium{{ID: "m1", UserID: env.User.ID, UserSpaceID: env.UserSpace.ID}},
					}
				},
			},
			content:   "Hoge dayo",
			MediumIDs: []string{"m1"},
		},
	}

	for i, test := range tests {
		// -------------------- preparation --------------------
		test.LogCase(t, i)
		api.CleanupDB()
		env := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"content":    test.content,
				"medium_ids": test.MediumIDs,
			},
		)
		env.SetupAuthorization(c)

		// -------------------- execution --------------------
		timelineH := NewTimeline()
		status, _, err := timelineH.CreatePost(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		// timeline post record created
		tps, err := reg.NewTimelinePostRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(tps))
		tp := tps[0]
		tp.UserID = env.User.ID
		tp.UserSpaceID = env.UserSpace.ID

		// thread created
		ts, err := reg.NewThreadRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ts))
		thread := ts[0]
		assert.Equal(t, env.UserSpace.ID, thread.UserSpaceID)

		tptrs, err := reg.NewTimelinePostThreadRelationRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(tptrs))
		assert.Equal(t, thread.ID, tptrs[0].ThreadID)
		assert.Equal(t, tp.ID, tptrs[0].TimelinePostID)

		// microp post created
		mps, err := reg.NewMicroPostRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(mps))
		mp := mps[0]
		assert.Equal(t, env.User.ID, mp.UserID)
		assert.Equal(t, env.UserSpace.ID, mp.UserSpaceID)
		assert.Equal(t, test.content, mp.Content)

		mpmrs, err := reg.NewMicroPostMediumRelationRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, len(test.MediumIDs), len(mpmrs))
		for i, mediumID := range test.MediumIDs {
			assert.Equal(t, mp.ID, mpmrs[i].MicroPostID)
			assert.Equal(t, mediumID, mpmrs[i].MediumID)
		}
	}
}
