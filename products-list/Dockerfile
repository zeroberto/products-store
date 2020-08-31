FROM gradle:6.6.0-jdk14 AS build

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN gradle build

FROM openjdk:14

COPY --from=build /usr/src/app/build/distributions/app.tgz ./

RUN tar -xzf /app.tgz && rm -rf /app.tgz

ENTRYPOINT [ "/app/bin/products-list" ]
