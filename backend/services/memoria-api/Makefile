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
	ENV=test reflex -r '(\.go|go\.mod)' -s -- go test ./infra/handler/

.PHONY: test-package-run
test-package-run:
	ENV=test DB_LOG_LEVEL=silent GIN_MODE=release \
			reflex -r '(\.go|go\.mod)' -s -- \
			go test -test.run AccountInvite ./infra/handler/

.PHONY: mockgen
mockgen:
	sh ./mockgen.sh
