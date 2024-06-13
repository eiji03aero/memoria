package registry

import (
	"testing"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/infra/tbl"
	"memoria-api/util"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestBeginCommit_S(t *testing.T) {
	// -------------------- preparation --------------------
	reg, err := NewBuilder().Build(BuilderBuildDTO{
		InitDB: util.BoolToPointer(true),
	})
	assert.NoError(t, err)

	forceDelete(reg.(*Registry).DB, "users")

	expUsers := []*tbl.User{
		{Name: "u1"},
	}

	// -------------------- execution --------------------
	reg.BeginTx()
	err = reg.NewUserRepository().Create(repository.UserCreateDTO{
		ID:            "u1",
		Name:          "u1",
		AccountStatus: "invited",
	})
	assert.NoError(t, err)
	reg.CommitTx()

	// -------------------- assertion --------------------
	users, err := reg.NewUserRepository().Find(&repository.FindOption{})
	assert.NoError(t, err)
	assert.Equal(t, len(expUsers), len(users))
	for i := range expUsers {
		assert.Equal(t, expUsers[i].Name, users[i].Name)
	}
}

func TestBeginCommit_F(t *testing.T) {
	// -------------------- preparation --------------------
	reg, err := NewBuilder().Build(BuilderBuildDTO{
		InitDB: util.BoolToPointer(true),
	})
	assert.NoError(t, err)

	forceDelete(reg.(*Registry).DB, "users")

	expUsers := []*tbl.User{}

	// -------------------- execution --------------------
	reg.BeginTx()
	err = reg.NewUserRepository().Create(repository.UserCreateDTO{
		ID:            "u1",
		Name:          "u1",
		AccountStatus: "invited",
	})
	assert.NoError(t, err)
	reg.RollbackTx()

	// -------------------- assertion --------------------
	users, err := reg.NewUserRepository().Find(&repository.FindOption{})
	assert.NoError(t, err)
	assert.Equal(t, len(expUsers), len(users))
}

func forceDelete(db *gorm.DB, tableName string) {
	prev := db.Logger
	db.Logger = prev.LogMode(logger.Silent)

	db.Exec("ALTER TABLE " + tableName + " DISABLE TRIGGER ALL")
	db.Exec("DELETE FROM " + tableName)
	db.Exec("ALTER TABLE " + tableName + " ENABLE TRIGGER ALL")

	db.Logger = prev
}
