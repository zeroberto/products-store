# products-list

This service provides a GRPC server for communication via HTTP2 with a method for listing discounted products.

## Architecture

* Ready to integrate a microservice ecosystem
* Written in Java 14
* Message transport via GRPC and HTTP2
* Communication with MongoDB database

## Building and running

This project makes use of profiles to define the execution settings. To run the application according to a specific profile, use the following command, informing the profile in the property **-Dprofile**.

> To run successfully, it is necessary to have a file named applicationProfileName.yml in the project resources folder.

In addition, this project is part of a larger project, which contains the complete execution via docker-compose.

To run it directly from this directory and make calls to the API, it is necessary to have the MongoDB server running.

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

The above command will run the application with the default values of the environment variables. These pre-existing values and variables can be viewed in the `.env` file.

Next, the execution command containing all available variables:

```shell script
PROFILE="prod" \
  PRODUCTS_LIST_OPTS="-DanyVariable=anyVariableValue" \
  PRODUCTS_MONGO_DB_USERNAME="root" \
  PRODUCTS_MONGO_DB_PASSWORD="insecure" \
  docker-compose up
```

### Local Gradle

#### Building

To build the project, run the command below:

```shell script
./gradlew clean build
```

#### Running

To execute the project, run the command below:

```shell script
./gradlew run
```

The default profile for running via gradle is empty, referencing the standard `application.yml` file. If you want to use another profile, inform the **-Dprofile** property after the command, as shown below:

```shell script
./gradlew run -Dprofile=local
```

> To run successfully, it is necessary to have a file named applicationProfileName.yml in the project resources folder.

### Tests

To run the tests, simply use the command:

```shell script
./gradlew test
```

#### Integration

To run the integration tests, simply use the command:

```shell script
./gradlew integTest
```

If you want to ensure that the projects used in the integration tests are in the latest versions of the repository, execute the following command. That's because the docker usually uses cache to make the build faster.

```shell script
cd src/integTest/resources/docker

docker-compose build --no-cache
```

#### Limitations

Integration tests are currently limited to Unix-based systems.

## Future improvements

* Tests to retrieve the values of the authentication settings with environment variables
* Create integration tests

## License

[MIT](LICENSE) License
