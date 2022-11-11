package constant

const CreateQuestionsQuery = `CREATE TABLE IF NOT EXISTS questions
(
    id SERIAL,
	user_id TEXT NOT NULL,
    tittle TEXT NOT NULL,
    statement TEXT NOT NULL,
    tags TEXT DEFAULT '',
	created_on TIMESTAMP NOT NULL,
    CONSTRAINT questions_pkey PRIMARY KEY (id)
)`
const CreateAnswersQuery = `CREATE TABLE IF NOT EXISTS answers
(
    id SERIAL,
	question_id INTEGER UNIQUE NOT NULL,
	user_id TEXT NOT NULL,
    comment TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL,
    CONSTRAINT answers_pkey PRIMARY KEY (id),
	CONSTRAINT fk_question_id FOREIGN KEY(question_id) REFERENCES questions(id) 
)`
