include .env
up:
	docker-compose --env-file .env up -d
# 	sleep 3
back:
	docker-compose down -v

down:
	docker-compose down

create-migrations:
	./migrate create -ext=sql -dir=internal/storage/pgstorage/migrations -seq init

migrate-up:
	./migrate -path=internal/storage/pgstorage/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate-down:
	./migrate -path=internal/storage/pgstorage/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
