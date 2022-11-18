package answer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/jtengananbd/questionsandanswers/constant"
)

func TestAnswerRepository_GetByQuestionID(t *testing.T) {
	// skip in short mode
	if testing.Short() {
		return
	}
	// Setup database
	dbContainer, connPool, err := SetupTestDatabase()
	if err != nil {
		t.Error(err)
	}
	defer dbContainer.Terminate(context.Background())

	ensureTableExists(db)
	clearTable(db)
	insertTable(db)

	// Create user repository
	repository := NewRepository(connPool)

	// Run tests against db
	t.Run("GetByQuestionID", func(t *testing.T) {
		answer, err := repository.GetByQuestionID("1")

		assert.NoError(t, err)
		assert.Equal(t, answer.ID, "1")
	})

}

func ensureTableExists(db *sql.DB) {
	if _, err := db.Exec(constant.CreateQuestionsQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(constant.CreateAnswersQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable(db *sql.DB) {
	db.Exec("DELETE FROM answers")
	db.Exec("ALTER SEQUENCE answers_id_seq RESTART WITH 1")
	db.Exec("DELETE FROM questions")
	db.Exec("ALTER SEQUENCE questions_id_seq RESTART WITH 1")
}
func insertTable(db *sql.DB) {
	db.Exec("INSERT INTO questions (user_id, tittle, statement, tags, created_on) VALUES ('julio@mail.com', 'tittle1', 'statement', 'GO, code', '2022-11-18 15:00:01'::timestamp)")
	db.Exec("INSERT INTO answers (question_id, user_id, comment, created_on) VALUES (1, 'julio@mail.com', 'comment', '2022-11-18 15:00:01'::timestamp)")
}

var db *sql.DB

func SetupTestDatabase() (testcontainers.Container, *sql.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "postgres",
			"POSTGRES_PASSWORD": "mypassword",
			"POSTGRES_USER":     "postgres",
		},
	}
	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, nil, err
	}
	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("the port : ", port)
	host, err := dbContainer.Host(context.Background())
	if err != nil {
		return nil, nil, err
	}
	user := "postgres"
	password := "mypassword"
	dbname := "postgres"

	connectionValues := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port.Port())

	db, err = sql.Open("postgres", connectionValues)
	if err != nil {
		return nil, nil, err
	}
	return dbContainer, db, err
}
