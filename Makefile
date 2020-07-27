repository = github.com/trueone
project = beetest

proto:
	mkdir -p /tmp/protogen/$(repository)/
	rm -rf /tmp/protogen/$(repository)/$(project)
	ln -s $(shell pwd) /tmp/protogen/$(repository)/$(project)
	find -iname '*.proto' -exec \
		protoc -I/usr/local/include -I. -I../ \
			--proto_path=. \
			--go_out=plugins=grpc:/tmp/protogen/ {} \
		\;

test:
	go test -v ./...

.PHONY: proto, test