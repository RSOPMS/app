include .env

.PHONY: $(wildcard *)

##  help; Print this help message
help:
	@echo Tip:
	@echo "  review Makefile for more detailed targets"
	@echo Usage:
	@sed -n 's/^##//p' Makefile | column -t -s ';' | sed -e 's/^//'

# =============================================================================
# Init
# =============================================================================

##  init; Initialize the repository for local development
init: clean
	@ln -s ../.env app-ingress/.env
	@ln -s ../.env app-bulk/.env
	@ln -s ../.env app-issue/.env
	@ln -s ../.env app-login/.env
	@ln -s ../.env app-static/.env
	@ln -s ../.env database/.env

	@ln -s ../.air.toml app-ingress/.air.toml
	@ln -s ../.air.toml app-bulk/.air.toml
	@ln -s ../.air.toml app-issue/.air.toml
	@ln -s ../.air.toml app-login/.air.toml
	@ln -s ../.air.toml app-static/.air.toml

##  clean; Remove all generated files
clean:
	@rm -rf app-bulk/.env
	@rm -rf app-ingress/.env
	@rm -rf app-issue/.env
	@rm -rf app-login/.env
	@rm -rf app-static/.env
	@rm -rf database/.env

	@rm -rf app-ingress/.air.toml
	@rm -rf app-bulk/.air.toml
	@rm -rf app-issue/.air.toml
	@rm -rf app-login/.air.toml
	@rm -rf app-static/.air.toml

# =============================================================================
# k8s dev
# =============================================================================

kd: k8s/dev
##  k8s/dev; Build and start the k8s cluster; (alias: kd)
k8s/dev: k8s/dev/build k8s/dev/start

k8s/dev/database: database/build database/start

k8s/dev/grafana: grafana/build grafana/start

k8s/dev/static: app-static/build app-static/start

k8s/dev/issue: app-issue/build app-issue/start

k8s/dev/bulk: app-bulk/build app-bulk/start

k8s/dev/ingress: app-ingress/build app-ingress/start

k8s/dev/login: app-login/build app-login/start

# -----------------------------------------------------------------------------
# k8s dev build
# -----------------------------------------------------------------------------

kdb: k8s/dev/build
##  k8s/dev/build; Build Docker container images; (alias: kdb)
k8s/dev/build: database/build grafana/build app-static/build app-issue/build app-bulk/build app-ingress/build app-login/build

database/build:
	@rm -rf ./database/initdb
	@mkdir -p ./database/initdb
	@cp ./database/alter/* ./database/initdb/
	@cp ./database/mock/* ./database/initdb/
	@docker build -f ./database/Dockerfile --tag bugbase-database:latest .

grafana/build:
	@echo "TODO"

app-static/build:
	@docker build -f ./app-static/Dockerfile --tag bugbase-static:latest .

app-issue/build:
	@docker build -f ./app-issue/Dockerfile --tag bugbase-issue:latest .

app-bulk/build:
	@docker build -f ./app-bulk/Dockerfile --tag bugbase-bulk:latest .

app-ingress/build:
	@docker build -f ./app-ingress/Dockerfile --tag bugbase-ingress:latest .

app-login/build:
	@docker build -f ./app-login/Dockerfile --tag bugbase-login:latest .

## ----------------------------------------------------------------------------
## k8s dev start
## ----------------------------------------------------------------------------

kds: k8s/dev/start
##  k8s/dev/start; Apply k8s configurations; (alias: kds)
k8s/dev/start: configmap/start secret/start ingress/start database/start nats/start grafana/start app-static/start app-issue/start app-bulk/start app-ingress/start app-login/start

configmap/start:
	@kubectl apply -f ./k8s/configmap.yaml

secret/start:
	@kubectl apply -f ./k8s/secret.yaml

ingress/start:
	@kubectl apply -f ./k8s/ingress.yaml

database/start:
	@kubectl apply -f ./k8s/database.yaml

grafana/start:
	@kubectl apply -f ./k8s/grafana.yaml

nats/start:
	@kubectl apply -f ./k8s/nats.yaml

app-static/start:
	@kubectl apply -f ./k8s/app-static.yaml

app-issue/start:
	@kubectl apply -f ./k8s/app-issue.yaml

app-bulk/start:
	@kubectl apply -f ./k8s/app-bulk.yaml

app-ingress/start:
	@kubectl apply -f ./k8s/app-ingress.yaml

app-login/start:
	@kubectl apply -f ./k8s/app-login.yaml

# -----------------------------------------------------------------------------
# k8s dev delete
# -----------------------------------------------------------------------------

kdd: k8s/dev/delete
##  k8s/dev/delete; Delete the k8s cluster; (alias: kdd)
k8s/dev/delete: configmap/delete secret/delete ingress/delete database/delete nats/delete grafana/delete app-static/delete app-issue/delete app-bulk/delete app-ingress/delete app-login/delete

configmap/delete:
	@kubectl delete -f ./k8s/configmap.yaml

secret/delete:
	@kubectl delete -f ./k8s/secret.yaml

ingress/delete:
	@kubectl delete -f ./k8s/ingress.yaml

database/delete:
	@kubectl delete -f ./k8s/database.yaml

grafana/delete:
	@kubectl delete -f ./k8s/grafana.yaml

nats/delete:
	@kubectl delete -f ./k8s/nats.yaml

app-static/delete:
	@kubectl delete -f ./k8s/app-static.yaml

app-issue/delete:
	@kubectl delete -f ./k8s/app-issue.yaml

app-bulk/delete:
	@kubectl delete -f ./k8s/app-bulk.yaml

app-ingress/delete:
	@kubectl delete -f ./k8s/app-ingress.yaml

app-login/delete:
	@kubectl delete -f ./k8s/app-login.yaml

# -----------------------------------------------------------------------------
# database
# -----------------------------------------------------------------------------

mu: migrate/up
##  migrate/up; Run up migrations; (alias: mu)
migrate/up:
	@goose -dir=database/migrations up

md: migrate/down
##  migrate/down; Run down migrations; (alias: md)
migrate/down:
	@goose -dir=database/migrations down

mf: migrate/fresh
##  migrate/fresh; Rebuild the database; (alias: mf)
migrate/fresh:
	@goose -dir=database/migrations reset
	@make migrate/up

mfs: migrate/fresh/seed
##  migrate/fresh/seed; Rebuild the database and seed it; (alias: mfs)
migrate/fresh/seed:
	@make migrate/fresh
	@goose -dir=database/seeds -no-versioning up
