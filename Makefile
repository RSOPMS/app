# =============================================================================
# Dev
# =============================================================================

.PHONY: k8s-dev
k8s-dev: k8s-dev-build k8s-dev-start

.PHONY: k8s-dev-static
k8s-dev: app-static-build app-static-start

.PHONY: k8s-dev-issue
k8s-dev: app-issue-build app-issue-start

## ----------------------------------------------------------------------------
## Dev build
## ----------------------------------------------------------------------------

.PHONY: k8s-dev-build
k8s-dev-build: app-static-build app-issue-build

.PHONY: app-static-build
app-static-build:
	@docker build -f ./app-static/Dockerfile --tag bugbase-static .

.PHONY: app-issue-build
app-issue-build:
	@docker build -f ./app-issue/Dockerfile --tag bugbase-issue .

## ----------------------------------------------------------------------------
## Dev start
## ----------------------------------------------------------------------------

.PHONY: k8s-dev-start
k8s-dev-start: ingress-start app-static-start app-issue-start

.PHONY: ingress-start
ingress-start:
	@kubectl apply -f ./k8s/ingress.yaml

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
k8s-dev-stop: ingress-stop app-static-stop app-issue-stop

.PHONY: ingress-stop
ingress-stop:
	@kubectl delete -f ./k8s/ingress.yaml

#.PHONY: database-stop
#database-stop:
#	@kubectl delete -f ./k8s/database.yaml

.PHONY: app-static-stop
app-static-stop:
	@kubectl delete -f ./k8s/app-static.yaml

.PHONY: app-issue-stop
app-issue-stop:
	@kubectl delete -f ./k8s/app-issue.yaml