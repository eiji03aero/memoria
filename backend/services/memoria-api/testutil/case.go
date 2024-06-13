package testutil

import (
	"strconv"
	"testing"

	"memoria-api/domain/cerrors"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TestCase struct {
	Name   string
	Seeder Seeder
}

type TestCaser interface {
	LogCase(t *testing.T, no int)
}

type Seeder func(env UserEnv) []any

func (c *TestCase) LogCase(t *testing.T, no int) {
	t.Log("Case " + strconv.Itoa(no) + ": " + c.Name)
}

type TestCaseAssertValidationErrorDTO struct {
	T        *testing.T
	ExpVKey  cerrors.ValidationKey
	ExpVName string
	Err      error
}

func (c *TestCase) AssertValidationError(dto TestCaseAssertValidationErrorDTO) {
	validationErr, ok := dto.Err.(cerrors.Validation)
	assert.True(dto.T, ok)

	assert.Equal(dto.T, dto.ExpVKey, validationErr.Key)
	assert.Equal(dto.T, dto.ExpVName, validationErr.Name)
}

func (c *TestCase) InstallSeeds(db *gorm.DB, env UserEnv) {
	seeds := c.Seeder(env)
	for _, seed := range seeds {
		db.Create(seed)
	}
}

func ExecuteTestCases[T TestCaser](
	t *testing.T,
	tests []T,
	fn func(test T),
) {
	for i, test := range tests {
		test.LogCase(t, i)

		fn(test)
	}
}
