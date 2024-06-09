package testutil

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CleanupDB(db *gorm.DB) {
	forceDelete(db, "users")
	forceDelete(db, "user_spaces")
	forceDelete(db, "user_user_space_relations")
	forceDelete(db, "user_invitations")
	forceDelete(db, "albums")
	forceDelete(db, "user_space_album_relations")
	forceDelete(db, "media")
	forceDelete(db, "album_medium_relations")
	forceDelete(db, "user_space_activities")
	forceDelete(db, "micro_post_medium_relations")
	forceDelete(db, "micro_posts")
	forceDelete(db, "thread_micro_post_relations")
	forceDelete(db, "threads")
	forceDelete(db, "timeline_post_thread_relations")
	forceDelete(db, "timeline_posts")
}

func forceDelete(db *gorm.DB, tableName string) {
	prev := db.Logger
	db.Logger = prev.LogMode(logger.Silent)

	db.Exec("ALTER TABLE " + tableName + " DISABLE TRIGGER ALL")
	db.Exec("DELETE FROM " + tableName)
	db.Exec("ALTER TABLE " + tableName + " ENABLE TRIGGER ALL")

	db.Logger = prev
}
