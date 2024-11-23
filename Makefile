DB_URL=postgres://thundio:password@postgres:5432/thundio_db?sslmode=disable


.PHONY: test sqlc clean

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

sqlc:
	@echo "Generating SQLC code..."
	@sqlc generate

test:
	@echo "Running Go Tests..."
	@go test ./... -v
