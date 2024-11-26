.PHONY: clean
clean:
	@rm -f db/database.db
	@cat alter/* | sqlite3 db/database.db

.PHONY: mock
mock: clean
	@cat mock/* | sqlite3 db/database.db

.PHONY: start
start:
	@rm -rf initdb
	@mkdir -p initdb
	@cp alter/* initdb/
	@cp mock/* initdb/
	@docker run -d --name bugbase-database --env-file=.env --mount type=bind,source=./initdb,target=/docker-entrypoint-initdb.d,readonly --publish 5432:5432 postgres

.PHONY: stop
stop:
	@docker stop bugbase-database
	@docker rm bugbase-database