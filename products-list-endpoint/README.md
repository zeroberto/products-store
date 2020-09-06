# products-list-endpoint

This service exposes a HTTP endpoint GET /product that returns a list of products.

## Architecture

* Ready to integrate a microservice ecosystem
* Written in GoLang
* Communication with other services via gRPC and HTTP2

## Building and running

This project makes use of profiles to define the execution settings. To run the application according to a specific profile, using any of the execution methods in the next sections, informing the profile in the go flag parameter **-fprofile**.

> To run successfully, it is necessary to have a file named applicationProfileName.yml in the project resources folder.

In addition, this project is part of a larger project, which contains the complete execution via docker-compose.

To run it directly from this directory and make calls to the API, it is necessary to have the [products-list](../products-list) application running.

### Docker

To run the application through a docker container, execute the following command inside the project's root folder.

```shell script
docker-compose up 
```

The default profile for running via the docker is **local**. If you want to use another profile, inform the **PROFILE** property before the command, as shown below:

```shell script
PROFILE=prod docker-compose up
```

> To run successfully, it is necessary to have a file named applicationProfileName.yml in the project resources folder.

### Local Build

#### Building

To build the project, run the command below:

```shell script
go build .
```

#### Running

To execute the project, run the command below:

```shell script
go run main.go
```

The default profile for running via gradle is empty, referencing the standard `application.yml` file. If you want to use another profile, inform the **-fprofile** property after the command, as shown below:

```shell script
go run main.go -fprofile=local
```

> To run successfully, it is necessary to have a file named applicationProfileName.yml in the project resources folder.

## Development

### Generating proto files

Installing compiler:

```shell script
sudo apt install -y protobuf-compiler
```

For more information, click [here](https://grpc.io/docs/protoc-installation/).

Running make routine:

```shell script
# GNU Make - Copyright (C) 1988-2014 Free Software Foundation, Inc.
make proto
```

## License

[MIT](LICENSE) License
