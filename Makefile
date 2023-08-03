tools:
	brew install mockery && \
	brew install golang-migrate && \
	brew install sqlc
mock:
	go generate ./...
sqlc:
	sqlc generate
db-schema:
	docker exec -it db pg_dump --schema-only postgres -f /db_schema/schema.sql -U postgres
migration:
	migrate create -ext sql -dir ./db/migrations -seq $(NAME)
migrate:
	migrate -path ./db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down
migrate-fix:
	migrate -path ./db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" force $(VERSION)
test:
	go test -v -cover ./...