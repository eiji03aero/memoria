.PHONY: build
build:
	go build -o ./bin/app ./cmd/app/main.go

.PHONY: run
run:
	reflex -r '(\.go|go\.mod)' -s go run ./cmd/app/*.go

# just a sample, do not run this
.PHONY: generate-migration
generate-migration:
	cd /backend
	echo migrate create -ext sql -dir db/migrations -seq create_users_table
