clean:
	rm -rf build

build:
	mkdir -p build

gen-proto:
	mkdir -p build/gen
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc \
  		--go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:build/gen \
  		--go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:build/gen \
  		protobuf/proto/productslist/products_list.proto
	mv build/gen/github.com/zeroberto/products-store/ pb/

proto: clean build gen-proto
