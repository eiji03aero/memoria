.PHONY: psql
psql:
	PGPASSWORD=${POSTGRES_PASSWORD} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

# just a sample, do not run this
.PHONY: generate-migration
generate-migration:
	echo migrate create -ext sql -dir db/migrations -seq create_users_table

.PHONY: migrate
migrate:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable' -verbose up

.PHONY: migrate-down-all
migrate-down-all:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable' -verbose down --all

.PHONY: db-reset
db-reset:
	make db-drop
	make migrate

.PHONY: db-drop
db-drop:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable' -verbose drop -f
