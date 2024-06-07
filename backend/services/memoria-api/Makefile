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
			go test -test.run MediumConfirmUploads -v ./infra/handler/

.PHONY: mockgen
mockgen:
	sh ./mockgen.sh

# === RUN   TestMediumConfirmUploads_S_Invoker
#     case.go:21: Case 0: Gets page of all
#     medium_test.go:551:
#                 Error Trace:    /backend/services/memoria-api/infra/handler/medium_test.go:551
#                 Error:          Not equal:
#                                 expected: "{\"UserID\": \"01HZSSQDKYHEWG53SH88N3RNER\",\"MediumIDs\":[\"m1\",\"m2\",\"m3\"]}"
#                                 actual  : "{\"user_id\": \"01HZSSQDKYHEWG53SH88N3RNER\", \"medium_ids\": [\"m1\", \"m2\", \"m3\"]}"
# 
#                                 Diff:
#                                 --- Expected
#                                 +++ Actual
#                                 @@ -1 +1 @@
#                                 -{"UserID": "01HZSSQDKYHEWG53SH88N3RNER","MediumIDs":["m1","m2","m3"]}
#                                 +{"user_id": "01HZSSQDKYHEWG53SH88N3RNER", "medium_ids": ["m1", "m2", "m3"]}
#                 Test:           TestMediumConfirmUploads_S_Invoker
# --- FAIL: TestMediumConfirmUploads_S_Invoker (0.15s)
