# syntax=docker/dockerfile:1

FROM golang:alpine3.19 AS dev
WORKDIR /app
COPY . .
ENV CGO_ENABLED 0
RUN go mod download

FROM dev as build-env
RUN go build -o dbMigrations cmd/dbMigrations/main.go

FROM alpine:3.14 as app
RUN useradd -m dbMigrations-user
USER dbMigrations-user
COPY --from=build-env /app/dbMigrations /usr/bin/dbMigrations

CMD ["/dbMigrations"]

#Just rebuild the golang code
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon


