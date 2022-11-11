# questionsandanswers

enable module in your go env
$ go env -w GO111MODULE=on 

$ go build

Start a PostgreSQL docker instance on port 5432
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mypassword -d postgres

$ go run main.go

You can fetch the Go modules using the following commands.

$ go get -u github.com/gorilla/mux 
$ go get -u github.com/lib/pq


Pending tasks:
    - docker conteinerization for both service and database for demo purposes 
    - Include validations for input objects, like required fields or email format with some validator lib
    - centralized error handling to unify error messages.
    - unit testing 