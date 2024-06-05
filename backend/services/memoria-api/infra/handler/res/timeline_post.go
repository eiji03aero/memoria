package res

import (
	"time"

	"memoria-api/domain/model"
)

type TimelinePost struct {
	ID        string    `json:"id"`
	Thread    *Thread   `json:"thread"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

func TimelinePostFromModel(m *model.TimelinePost) (tp *TimelinePost) {
	tp = &TimelinePost{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
	}

	if m.Thread != nil {
		tp.Thread = ThreadFromModel(m.Thread)
	}

	if m.User != nil {
		tp.User = UserFromModel(m.User)
	}

	return
}
