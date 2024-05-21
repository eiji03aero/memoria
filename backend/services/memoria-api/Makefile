.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux go build -o /memoria-api ./cmd/app/*.go

.PHONY: run
run:
	reflex -r '(\.go|go\.mod)' -s go run ./cmd/app/*.go

.PHONY: test
test:
	ENV=test go test -v ./...

.PHONY: test-package
test-package:
	ENV=test reflex -r '(\.go|go\.mod)' -s go test -v ./infra/handler/

.PHONY: mockgen
mockgen:
	sh ./mockgen.sh
