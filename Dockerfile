FROM golang:1.21 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY api ./api
COPY cmd ./cmd
COPY database ./database
COPY transaction ./transaction
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin ./...

FROM alpine:latest AS deploy

COPY --from=build /usr/local/bin/pseudo-bank /usr/local/bin/pseudo-bank
