package answer

import (
	"context"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/integration"
)

func TestAnswerRepository_GetByQuestionIDIntegration(t *testing.T) {
	// skip in short mode
	if testing.Short() {
		return
	}
	// Setup database
	dbContainer, db, err := integration.SetupTestDatabase()
	if err != nil {
		t.Error(err)
	}
	defer dbContainer.Terminate(context.Background())

	integration.EnsureTableExists(db)
	integration.ClearTable(db)
	integration.InsertTable(db)

	// Create user repository
	repository := NewRepository(db)

	// Run tests against db
	t.Run("GetByQuestionID", func(t *testing.T) {
		answer, err := repository.GetByQuestionID("1")

		assert.NoError(t, err)
		assert.Equal(t, answer.ID, "1")
	})

}
func TestAnswerRepository_UpdateIntegration(t *testing.T) {
	// skip in short mode
	if testing.Short() {
		return
	}
	// Setup database
	dbContainer, db, err := integration.SetupTestDatabase()
	if err != nil {
		t.Error(err)
	}
	defer dbContainer.Terminate(context.Background())

	integration.EnsureTableExists(db)
	integration.ClearTable(db)
	integration.InsertTable(db)

	// Create user repository
	repository := NewRepository(db)

	// Run tests against db
	t.Run("Update", func(t *testing.T) {
		answerDB, err := repository.GetByQuestionID("1")
		if err != nil {
			t.Error(err)
		}
		answerDB.Comment = "Comment updated"

		answerRs, err := repository.Update(answerDB)

		assert.NoError(t, err)
		assert.Equal(t, answerRs.Comment, "Comment updated")
	})

}
func TestAnswerRepository_CreateIntegration(t *testing.T) {
	// skip in short mode
	if testing.Short() {
		return
	}
	// Setup database
	dbContainer, db, err := integration.SetupTestDatabase()
	if err != nil {
		t.Error(err)
	}
	defer dbContainer.Terminate(context.Background())

	integration.EnsureTableExists(db)
	integration.ClearTable(db)
	integration.InsertTable(db)

	// Create user repository
	repository := NewRepository(db)

	answer := entity.Answer{
		QuestionID: 2,
		UserID:     "user@mail.com",
		Comment:    "comment test",
	}

	// Run tests against db
	t.Run("Create", func(t *testing.T) {
		answerRs, err := repository.Create(answer)

		assert.NoError(t, err)
		assert.Equal(t, answerRs.Comment, "comment test")
	})

}

func TestAnswerRepository_DeleteIntegration(t *testing.T) {
	// skip in short mode
	if testing.Short() {
		return
	}
	// Setup database
	dbContainer, db, err := integration.SetupTestDatabase()
	if err != nil {
		t.Error(err)
	}
	defer dbContainer.Terminate(context.Background())

	integration.EnsureTableExists(db)
	integration.ClearTable(db)
	integration.InsertTable(db)

	// Create user repository
	repository := NewRepository(db)

	// Run tests against db
	t.Run("Delete", func(t *testing.T) {
		err := repository.DeleteByQuestionID("1")
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, err)

		answerRs, err := repository.GetByQuestionID("1")
		assert.Error(t, err)
		assert.Equal(t, "Resource Answer with ID:1 not found", err.Error())
		assert.Empty(t, answerRs.ID)
	})

}
