# -------------------------------------
# BUILDER
# -------------------------------------

FROM golang:alpine AS build

RUN apk add --no-cache git make build-base

WORKDIR /app

COPY . .

WORKDIR /app/app-issue

ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

RUN go mod download

RUN go build -ldflags "-s -w -extldflags '-static'" -o ./bin/service ./cmd/main.go

# -------------------------------------
# RUNNER
# -------------------------------------

FROM scratch

ENV APP_PORT=8080 DB_PATH="./db/database.db"

WORKDIR /

COPY --from=build /app/app-issue/bin/service /service

EXPOSE ${APP_PORT}

CMD [ "/service" ]