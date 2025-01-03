# -------------------------------------
# BUILDER
# -------------------------------------

FROM golang:alpine AS build

WORKDIR /app

COPY . .

WORKDIR /app/app-ingress

RUN go mod download

RUN go build -o ./bin/service ./cmd/main.go

# -------------------------------------
# RUNNER
# -------------------------------------

FROM scratch

ENV APP_PORT=8080

WORKDIR /

COPY --from=build /app/app-ingress/bin/service /service

EXPOSE ${APP_PORT}

CMD [ "/service" ]
