package dao

import (
	"sort"

	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/domain/value"
	"memoria-api/infra/tbl"
)

type timeline[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewTimeline(db *gorm.DB) repository.Timeline {
	return &timeline[tbl.TimelineUnit]{db: db}
}

func (d *timeline[T]) Find(fOpt *repository.FindOption) (tus model.TimelineUnits, cpagi repository.CPagination, err error) {
	tpRepo := NewTimelinePost(d.db)
	usaRepo := NewUserSpaceActivity(d.db)
	unitTbls := []*tbl.TimelineUnit{}

	// -------------------- build timeline_units --------------------
	tpsQ := d.db.
		Select("id", "'timeline-post' AS type", "'{}'::jsonb AS data", "created_at").
		Model(&tbl.TimelinePost{}).
		Where("user_space_id = ?", fOpt.Filter["user_space_id"])

	usasQ := d.db.
		Select("id", "'user-space-activity' AS type", "data::jsonb AS data", "created_at").
		Model(&tbl.UserSpaceActivity{}).
		Where("user_space_id = ?", fOpt.Filter["user_space_id"])

	timelineUnitsQ := d.db.Raw("? UNION ?", tpsQ, usasQ)

	// -------------------- build main query --------------------
	mainQ := d.db.Select("*").Table("timeline_units")
	shouldReverse := false
	if fOpt.Cursor == "" {
		mainQ = mainQ.Limit(100).Order("id DESC")
	}
	if fOpt.Cursor != "" {
		if fOpt.CBefore != 0 {
			shouldReverse = true
			mainQ = mainQ.Where("id > ?", fOpt.Cursor).Limit(fOpt.CBefore).Order("id ASC")
		} else if fOpt.CAfter != 0 {
			mainQ = mainQ.Where("id < ?", fOpt.Cursor).Limit(fOpt.CAfter).Order("id DESC")
		}

		if !fOpt.CExclude {
			mainQ = mainQ.Or("id = ?", fOpt.Cursor)
		}
	}

	err = d.db.Raw("WITH timeline_units AS (?) ?", timelineUnitsQ, mainQ).Scan(&unitTbls).Error
	if err != nil {
		return
	}

	// -------------------- build domain models --------------------
	tus = model.TimelineUnits{}
	for _, unitTbl := range unitTbls {
		tu, e := unitTbl.ToModel()
		if e != nil {
			err = e
			return
		}

		tus = append(tus, tu)
	}

	for _, tu := range tus {
		if tu.Type == value.TimelineUnitType_TimelinePost {
			tp, e := tpRepo.FindOneByID(tu.ID, &repository.FindOption{
				Preloads: []string{"Thread.MicroPosts.Media", "User"},
			})
			if e != nil {
				err = e
				return
			}

			tu.Data = tp
		}

		if tu.Type == value.TimelineUnitType_UserSpaceActivity {
			usa, e := usaRepo.FindOneByID(tu.ID)
			if e != nil {
				err = e
				return
			}

			tu.Data = usa
		}
	}

	if shouldReverse {
		sort.Sort(sort.Reverse(tus))
	}

	// -------------------- sets cursors --------------------
	prevCursor := func() string {
		if len(tus) == 0 {
			return fOpt.Cursor
		}

		return tus[0].ID
	}()

	nextCursor := func() string {
		if len(tus) == 0 {
			return fOpt.Cursor
		}

		return tus[len(tus)-1].ID
	}()

	cpagi = repository.CPagination{
		PrevCursor: prevCursor,
		NextCursor: nextCursor,
	}

	return
}
