# Makefile

# Varsayılan hedefler
.PHONY: test sqlc clean

# SQLC kodlarını oluşturma
sqlc:
	@echo "Generating SQLC code..."
	@sqlc generate

# Test komutu
test:
	@echo "Running Go Tests..."
	@go test ./... -v
