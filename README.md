## Bootcamp project questionsandanswers


### Requirements
To run the project you need installed:
    * docker
    * docker-compose

#### To Build 
Execute this commnad to build the project into a docker compose

```$ docker-compose build```
#### To Run 
Execute this command to run the project, It will run on your local machine on port 8080, also the Database dependency will be started.

```$ docker-compose up```

#### To Play with APIs
There is a Postman collection under /postman_collection folder, just import and execute them. 


#### Execute tests
The tests are Unit tests and work as examples of how to implements test for the diffferent layers.
There are test for Repository, Service and APIs
to run all the test, exetute this command:
```$ go test ./...```   


Execute just the Unit tests
```$ go test -short ./...```


Execute just the Integration tests, It find your tests with name ending with Integration, e.g ```TestAnswerRepository_CreateIntegration```

```$ go test -v -run 'Integration' ./...```




## Pending tasks:
    * Refactor to use context package
    * Include validations for input objects, like required fields or email format with some validator lib
    * Centralized error handling to unify error messages.
    * Add unit test coverage using the tests examples developed for each layer.





#### old non used commands

to create an image using Dockerfile

```$ docker build --tag questions-and-answers .```


enable module in your go env

```$ go env -w GO111MODULE=on ```


```$ go build```

Start a PostgreSQL docker instance on port 5432

```docker run -p 5432:5432 -e POSTGRES_PASSWORD=mypassword -d postgres```


```$ go run main.go```

You can fetch the Go modules using the following commands.

```$ go get -u github.com/gorilla/mux ```

```$ go get -u github.com/lib/pq```

