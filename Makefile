tools:
	brew install mockery && \
	brew install golang-migrate
gen:
	go generate ./...
migration:
	migrate create -ext sql -dir ./db/migrations -seq $(NAME)
migrate:
	migrate -path ./db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down
migrate-fix:
	migrate -path ./db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" force $(VERSION)