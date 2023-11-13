FROM golang:1.21 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY api ./api
COPY cmd ./cmd
COPY database ./database
COPY transaction ./transaction
COPY user ./user
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin ./...

FROM alpine:latest AS deploy
 
# Add curl to allow for service healthcheck to be run
RUN apk update &&\
    apk upgrade &&\
    apk add curl

COPY --from=build /usr/local/bin/pseudo-bank /usr/local/bin/pseudo-bank
