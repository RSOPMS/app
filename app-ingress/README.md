# Ingress

Microservice for inserting (bulk) data into a database

## Development

Run the following command to start the development server

```sh
make dev
```

Test and vet the code using

```sh
make test
make vet
```

Build the project binary

```sh
make all
```

## NATS (local development)

Pull NATS Docker image
```sh
docker pull nats
```

Run NATS Server in Docker container
```sh
docker run -d --name nats-server -p 4222:4222 -p 8222:8222 nats
```

Stop NATS server
```sh
docker stop nats-server
```

Remove NATS server
```sh
docker rm nats-server
```
