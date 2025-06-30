include .env

# for Windows only
# Database connection string
DB_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${DB_NAME}?sslmode=disable

.PHONY: help migrate-create migrate-up migrate-down migrate-force migrate-version

# Show help
help:
	@echo "Available commands:"
	@echo "  migrate-create [name] 		Create a new migration file"
	@echo "  migrate-up            		Run all pending migrations"
	@echo "  migrate-down          		Rollback last migration"
	@echo "  migrate-version       		Show current migration version"
	@echo "  migrate-force         		Force set migration version"

# Create new migration files
migrate-create:
	@if "$(name)"=="" ( \
		echo Usage: make migrate-create name=your_migration_name & \
		exit 1 \
	) else ( \
		migrate create -ext sql -dir migrations -seq $(name) \
	)


# Run all pending migrations
migrate-up:
	migrate -path migrations -database "${DB_URL}" -verbose up

# Rollback last migration
migrate-down:
	migrate -path migrations -database "${DB_URL}" -verbose down 1

# Show current migration version
migrate-version:
	migrate -path migrations -database "${DB_URL}" version

# Force set migration version (useful for troubleshooting)
migrate-force:
	@echo "Enter version to force:" && \
	set /p version= && \
	migrate -path migrations -database "${DB_URL}" force $$version