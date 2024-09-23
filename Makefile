include .env

create_migration:
	migrate create -ext=sql -dir=migrations -tz "Asia/Jakarta" $(name)

migrate_up:
	migrate -path=migrations -database "${DATABASE_URL}" -verbose up

migrate_down:
	migrate -path=migrations -database "${DATABASE_URL}" -verbose down

.PHONY: create_migration migrate_up migrate_down