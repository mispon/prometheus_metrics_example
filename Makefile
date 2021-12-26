.PHONY: deps
deps:
	ls go.mod || go mod init
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: build
build: .generate .build

.PHONY: .generate
.generate:
	mkdir -p pkg/metrics
	protoc \
		   --go_out=pkg/metrics --go_opt=paths=import \
		   --go-grpc_out=pkg/metrics --go-grpc_opt=paths=import \
		   api/metrics.proto

.PHONY: .build
.build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/app cmd/main.go