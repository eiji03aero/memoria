package testutil

import "gorm.io/gorm"

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
}

func forceDelete(db *gorm.DB, tableName string) {
	db.Exec("ALTER TABLE " + tableName + " DISABLE TRIGGER ALL")
	db.Exec("DELETE FROM " + tableName)
	db.Exec("ALTER TABLE " + tableName + " ENABLE TRIGGER ALL")
}
