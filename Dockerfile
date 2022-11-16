# https://docs.docker.com/language/golang/build-images/

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o ./questionsandanswers

EXPOSE 8080

CMD ["./questionsandanswers"]