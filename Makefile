migrate:
	migrate -source file://db/migrations -database sqlite3://db/dev.db up
rollback:
	migrate -source file://db/migrations -database sqlite3://db/dev.db down 1
run:
	go run main.go