.PHONY: start
start:
	@rm -rf initdb
	@mkdir -p initdb
	@cp alter/* initdb/
	@docker run -d --name bugbase-database --env-file=.env --mount type=bind,source=./initdb,target=/docker-entrypoint-initdb.d,readonly --publish 5432:5432 postgres

.PHONY: start-mock
start-mock:
	@rm -rf initdb
	@mkdir -p initdb
	@cp alter/* initdb/
	@cp mock/* initdb/
	@docker run -d --name bugbase-database --env-file=.env --mount type=bind,source=./initdb,target=/docker-entrypoint-initdb.d,readonly --publish 5432:5432 postgres

.PHONY: stop
stop:
	@docker stop bugbase-database
	@docker rm bugbase-database