package res

import (
	"time"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/model"
	"memoria-api/domain/value"
)

type TimelineUnit struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Data      any       `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func TimelineUnitFromModel(m *model.TimelineUnit) (tu *TimelineUnit) {
	data := func() any {
		if m.Type == value.TimelineUnitType_UserSpaceActivity {
			usa, ok := m.Data.(*model.UserSpaceActivity)
			if !ok {
				panic(cerrors.NewInternal("failed to parse timeline unit in res"))
			}

			usaRes, err := UserSpaceActivityFromModel(usa)
			if err != nil {
				panic(err)
			}

			return usaRes
		}

		if m.Type == value.TimelineUnitType_TimelinePost {
			tp, ok := m.Data.(*model.TimelinePost)
			if !ok {
				panic(cerrors.NewInternal("failed to parse timeline unit in res"))
			}

			return TimelinePostFromModel(tp)
		}

		panic(cerrors.NewNotImplemented("unknown type for timeline unit: " + string(m.Type)))
	}()

	tu = &TimelineUnit{
		ID:        m.ID,
		Type:      string(m.Type),
		Data:      data,
		CreatedAt: m.CreatedAt,
	}
	return
}
