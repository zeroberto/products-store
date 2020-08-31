clean:
	rm -rf build

build:
	mkdir -p build

gen-proto:
	mkdir -p build/gen
	go get github.com/golang/protobuf/protoc-gen-go
	export PATH="$$PATH:$$(go env GOPATH)/bin"
	protoc \
  		--go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:build/gen \
  		--go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:build/gen \
  		protobuf/proto/productslist/products_list.proto
	rm -rf pb
	mv build/gen/github.com/zeroberto/products-store/ pb/
	rm -rf build

proto: clean build gen-proto
