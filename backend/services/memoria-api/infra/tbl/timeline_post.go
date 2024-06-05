package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type TimelinePost struct {
	ID          string `gorm:"column:id"`
	UserID      string `gorm:"column:user_id"`
	UserSpaceID string `gorm:"column:user_space_id"`
	User        *User
	Thread      []*Thread `gorm:"many2many:timeline_post_thread_relations;foreignKey:ID;joinForeignKey:TimelinePostID;References:ID;joinReferences:ThreadID"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t TimelinePost) TableName() string {
	return "timeline_posts"
}

func (t TimelinePost) ToModel() (m *model.TimelinePost, err error) {
	user, err := func() (u *model.User, e error) {
		if t.User == nil {
			return
		}

		u, e = t.User.ToModel()
		return
	}()
	if err != nil {
		return
	}

	thread := func() (th *model.Thread) {
		if t.Thread == nil || len(t.Thread) == 0 {
			return
		}

		th = t.Thread[0].ToModel()
		return
	}()

	m = model.NewTimelinePost(model.NewTimelinePostDTO{
		ID:          t.ID,
		UserID:      t.UserID,
		UserSpaceID: t.UserSpaceID,
		User:        user,
		Thread:      thread,
	})

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}

func (t *TimelinePost) FromModel(m *model.TimelinePost) {
	t.ID = m.ID
	t.UserID = m.UserID
	t.UserSpaceID = m.UserSpaceID
	t.CreatedAt = m.CreatedAt
	t.UpdatedAt = m.UpdatedAt
}
