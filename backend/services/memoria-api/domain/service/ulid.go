package service

import "github.com/oklog/ulid/v2"

func GenerateUlid() string {
	return ulid.Make().String()
}
