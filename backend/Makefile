# just a sample, do not run this
.PHONY: generate-migration
generate-migration:
	migrate create -ext sql -dir db/migrations -seq create_users_table

.PHONY: db-migrate
db-migrate:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable' -verbose up


.PHONY: db-down-1
db-down-1:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable' -verbose down 1

.PHONY: db-reset
db-reset:
	make db-drop
	make db-migrate

.PHONY: db-drop
db-drop:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable' -verbose drop -f

.PHONY: test-db-setup
test-db-setup:
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/memoria-test?sslmode=disable' -verbose drop -f
	migrate --path db/migrations --database 'postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/memoria-test?sslmode=disable' -verbose up
