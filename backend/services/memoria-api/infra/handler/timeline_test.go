package handler

import (
	"net/http"
	"testing"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/testutil"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

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
			assert.Equal(t, tp.ID, mpmrs[i].MicroPostID)
			assert.Equal(t, mediumID, mpmrs[i].MediumID)
		}
	}
}
