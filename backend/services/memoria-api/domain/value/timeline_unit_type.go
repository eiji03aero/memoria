package value

import (
	"fmt"

	"memoria-api/domain/cerrors"
)

type TimelineUnitType string

var (
	TimelineUnitType_UserSpaceActivity = TimelineUnitType("user-space-activity")
	TimelineUnitType_Post              = TimelineUnitType("post")
	TimelineUnitTypes                  = []TimelineUnitType{
		TimelineUnitType_UserSpaceActivity,
		TimelineUnitType_Post,
	}
)

func IsValidTimelineUnitType(t string) bool {
	for _, tut := range TimelineUnitTypes {
		if string(tut) == t {
			return true
		}
	}
	return false
}

func NewTimelineUnitType(t string) (tut TimelineUnitType, err error) {
	if !IsValidTimelineUnitType(t) {
		err = cerrors.NewConsistency(fmt.Sprintf("invalid timeline unit type: %s", t))
		return
	}

	tut = TimelineUnitType(t)
	return
}
