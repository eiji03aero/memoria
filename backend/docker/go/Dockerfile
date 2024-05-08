FROM golang:1.22.2-alpine3.19

RUN apk add bash make

RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1 \
  && go install github.com/cespare/reflex@v0.3.1 \
  && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1 \
  && go install go.uber.org/mock/mockgen@latest

CMD ["/bin/bash"]
