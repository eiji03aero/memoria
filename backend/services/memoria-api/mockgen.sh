mockgen -source=domain/interfaces/mailer.go -destination=testutil/mock/mailer.go -package mock
mockgen -source=domain/interfaces/s3.go -destination=testutil/mock/s3.go -package mock
mockgen -source=domain/interfaces/bgjob.go -destination=testutil/mock/bgjob.go -package mock
mockgen -source=domain/interfaces/registry.go -destination=testutil/mock/registry.go -package mock
