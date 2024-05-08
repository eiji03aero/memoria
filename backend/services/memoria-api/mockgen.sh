mockgen -source=domain/interfaces/mailer.go -destination=testutil/mock/mailer.go -package mock
mockgen -source=registry/registry.go -destination=testutil/mock/registry.go -package mock
