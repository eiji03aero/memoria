.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux go build -o /memoria-api ./cmd/app/*.go

.PHONY: run
run:
	reflex -r '(\.go|go\.mod)' -s go run ./cmd/app/*.go

.PHONY: test
test:
	ENV=test GIN_MODE=release DB_LOG_LEVEL=silent \
			go test -v ./...

.PHONY: test-run
test-run:
	ENV=test GIN_MODE=release DB_LOG_LEVEL=info \
			reflex -r '(\.go|go\.mod)' -s -- \
			go test -test.run Test -v ./infra/registry/

.PHONY: mockgen
mockgen:
	sh ./mockgen.sh
