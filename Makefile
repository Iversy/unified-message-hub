include .env

copy:
	cp env.cfg .env

up:
	docker-compose --env-file .env up -d
back:
	docker-compose down -v

down:
	docker-compose down

init-citus:
	@echo "Adding worker nodes to the Citus coordinator..."
	docker exec -it $(COORDINATOR) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c "SELECT master_add_node('$(WORKER1)', ${POSTGRES_PORT});"
	docker exec -it $(COORDINATOR) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c "SELECT master_add_node('$(WORKER2)', ${POSTGRES_PORT});"
	@echo "Worker nodes added successfully."

create-migrations:
	./migrate create -ext=sql -dir=internal/storage/pgstorage/migrations -seq init

migrate-up:
	./migrate -path=internal/storage/pgstorage/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate-down:
	./migrate -path=internal/storage/pgstorage/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
