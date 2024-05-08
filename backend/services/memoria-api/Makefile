.PHONY: build
build:
	go build -o ./bin/app ./cmd/app/main.go

.PHONY: run
run:
	reflex -r '(\.go|go\.mod)' -s go run ./cmd/app/*.go

.PHONY: test
test:
	ENV=test go test -v ./...

.PHONY: test-package
test-package:
	ENV=test reflex -r '(\.go|go\.mod)' -s go test -v ./

.PHONY: mockgen
mockgen:
	sh ./mockgen.sh
