module github.com/zeroberto/products-store/products-list-endpoint

go 1.15

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.33.0-dev
	google.golang.org/protobuf v1.25.0
)

replace github.com/golang/protobuf/protoc-gen-go => github.com/golang/protobuf/protoc-gen-go v1.4.2
replace google.golang.org/grpc => google.golang.org/grpc v1.33.0-dev
