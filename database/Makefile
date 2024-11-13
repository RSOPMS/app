clean:
	@rm -f db/database.db
	@cat alter/* | sqlite3 db/database.db

mock: clean
	@cat mock/* | sqlite3 db/database.db