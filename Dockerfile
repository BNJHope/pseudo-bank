FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
COPY database ./database
COPY transaction ./transaction
RUN go build -v -o /usr/local/bin ./...
