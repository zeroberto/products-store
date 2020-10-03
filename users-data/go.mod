module github.com/zeroberto/products-store/users-data

go 1.15

require (
	github.com/golang/protobuf v1.4.2
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/zeroberto/integration-test-suite v0.0.0-20201003165752-fcdbb5c41a88
	google.golang.org/grpc v1.33.0-dev
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/golang/protobuf/protoc-gen-go => github.com/golang/protobuf/protoc-gen-go v1.4.2

replace google.golang.org/grpc => google.golang.org/grpc v1.33.0-dev
