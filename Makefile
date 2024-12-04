# =============================================================================
# Dev
# =============================================================================

.PHONY: k8s-dev
k8s-dev: k8s-dev-build k8s-dev-start

.PHONY: k8s-dev-database
k8s-dev-database: database-build database-start

.PHONY: k8s-dev-static
k8s-dev-static: app-static-build app-static-start

.PHONY: k8s-dev-issue
k8s-dev-issue: app-issue-build app-issue-start

## ----------------------------------------------------------------------------
## Dev build
## ----------------------------------------------------------------------------

.PHONY: k8s-dev-build
k8s-dev-build: database-build app-static-build app-issue-build

.PHONY: database-build
database-build:
	@rm -rf ./database/initdb
	@mkdir -p ./database/initdb
	@cp ./database/alter/* ./database/initdb/
	@cp ./database/mock/* ./database/initdb/
	@docker build -f ./database/Dockerfile --tag bugbase-database:latest .

.PHONY: app-static-build
app-static-build:
	@docker build -f ./app-static/Dockerfile --tag bugbase-static:latest .

.PHONY: app-issue-build
app-issue-build:
	@docker build -f ./app-issue/Dockerfile --tag bugbase-issue:latest .

## ----------------------------------------------------------------------------
## Dev start
## ----------------------------------------------------------------------------

.PHONY: k8s-dev-start
k8s-dev-start: configmap-start secret-start ingress-start database-start app-static-start app-issue-start

.PHONY: configmap-start
configmap-start:
	@kubectl apply -f ./k8s/configmap.yaml

.PHONY: secret-start
secret-start:
	@kubectl apply -f ./k8s/secret.yaml

.PHONY: ingress-start
ingress-start:
	@kubectl apply -f ./k8s/ingress.yaml

.PHONY: database-start
database-start:
	@kubectl apply -f ./k8s/database.yaml

.PHONY: app-static-start
app-static-start:
	@kubectl apply -f ./k8s/app-static.yaml

.PHONY: app-issue-start
app-issue-start:
	@kubectl apply -f ./k8s/app-issue.yaml

## ----------------------------------------------------------------------------
## Dev stop
## ----------------------------------------------------------------------------

.PHONY: k8s-dev-stop
k8s-dev-stop: configmap-stop secret-stop ingress-stop database-stop app-static-stop app-issue-stop

.PHONY: configmap-stop
configmap-stop:
	@kubectl delete -f ./k8s/configmap.yaml

.PHONY: secret-stop
secret-stop:
	@kubectl delete -f ./k8s/secret.yaml

.PHONY: ingress-stop
ingress-stop:
	@kubectl delete -f ./k8s/ingress.yaml

.PHONY: database-stop
database-stop:
	@kubectl delete -f ./k8s/database.yaml

.PHONY: app-static-stop
app-static-stop:
	@kubectl delete -f ./k8s/app-static.yaml

.PHONY: app-issue-stop
app-issue-stop:
	@kubectl delete -f ./k8s/app-issue.yaml