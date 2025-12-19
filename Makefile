include .env
up:
	docker-compose up -d
# 	sleep 3
back:
	docker-compose down -v

down:
	docker-compose down

create_migration:
	./migrate create -ext=sql -dir=internal/storage/pgstorage/migrations -seq init

migrate-up:
	./migrate -path=internal/storage/pgstorage/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate-down:
	./migrate -path=internal/storage/pgstorage/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
