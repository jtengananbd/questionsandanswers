package question

import (
	"context"
	"testing"

	_ "github.com/lib/pq"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/integration"
	"github.com/stretchr/testify/assert"
)

func TestQuestionRepository_GetByIDIntegration(t *testing.T) {
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

	repository := NewRepository(db)

	// Run tests against db
	t.Run("GetByID", func(t *testing.T) {
		question, err := repository.GetByID("1")

		assert.NoError(t, err)
		assert.Equal(t, question.ID, "1")
	})

}

func TestQuestionRepository_CreateIntegration(t *testing.T) {
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

	repository := NewRepository(db)

	question := entity.Question{
		UserID:    &userID,
		Tittle:    "tittle test",
		Statement: "how to avoid ...",
		Tags:      "code, JS",
	}

	t.Run("Create", func(t *testing.T) {
		questionRs, err := repository.Create(question)

		assert.NoError(t, err)
		assert.Equal(t, questionRs.Statement, "how to avoid ...")
	})

}

func TestQuestionRepository_ListIntegration(t *testing.T) {
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

	repository := NewRepository(db)

	t.Run("List", func(t *testing.T) {
		questions, err := repository.List("julio@mail.com")

		assert.NoError(t, err)
		assert.NotEmpty(t, questions)
		assert.Equal(t, questions[0].ID, "1")
	})

}

func TestQuestionRepository_UpdateIntegration(t *testing.T) {
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

	repository := NewRepository(db)

	t.Run("Update", func(t *testing.T) {
		questionDB, err := repository.GetByID("1")
		if err != nil {
			t.Error(err)
		}
		questionDB.Statement = "Statement updated"

		questionRs, err := repository.Update(questionDB)

		assert.NoError(t, err)
		assert.Equal(t, questionRs.Statement, "Statement updated")
	})

}

func TestQuestionRepository_DeleteIntegration(t *testing.T) {
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

	repository := NewRepository(db)

	t.Run("Delete", func(t *testing.T) {
		err := repository.Delete("3")
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, err)

		questionRs, err := repository.GetByID("3")
		assert.Error(t, err)
		assert.Equal(t, "Resource Question with ID:3 not found", err.Error())
		assert.Empty(t, questionRs.ID)
	})

}
