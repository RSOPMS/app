.PHONY: start
start:
	@rm -rf initdb
	@mkdir -p initdb
	@cp alter/* initdb/
	@docker build --file Dockerfile.local --tag bugbase-database:latest .
	@docker run --detach --name bugbase-database --publish 5432:5432 bugbase-database

.PHONY: start-mock
start-mock:
	@rm -rf initdb
	@mkdir -p initdb
	@cp alter/* initdb/
	@cp mock/* initdb/
	@docker build --file Dockerfile.local --tag bugbase-database:latest .
	@docker run --detach --name bugbase-database --publish 5432:5432 bugbase-database

.PHONY: stop
stop:
	@docker stop bugbase-database
	@docker rm bugbase-database