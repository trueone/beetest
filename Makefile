.PHONY: proto
proto:
	protoc -I/usr/local/include -I. -I../ \
      -I${GOPATH}/src \
      --go_out=plugins=grpc:. \
      internal/app/secret/presenter/http/grpc/proto/secret.proto

.PHONY: test
test:
	go test -v ./...