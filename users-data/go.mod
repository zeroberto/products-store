module github.com/zeroberto/products-store/users-data

go 1.15

require (
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.33.0-dev
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	grpc.go4.org v0.0.0-20170609214715-11d0a25b4919
)

replace github.com/golang/protobuf/protoc-gen-go => github.com/golang/protobuf/protoc-gen-go v1.4.2

replace google.golang.org/grpc => google.golang.org/grpc v1.33.0-dev
