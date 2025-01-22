FROM golang

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app

RUN go build -o main ./cmd/server/main.go

EXPOSE 8080
CMD [ "/app/main" ]