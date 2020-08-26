FROM gradle:6.6.0-jdk14 AS build

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN gradle build

FROM openjdk:14

MAINTAINER zeroberto

COPY --from=build /usr/src/app/build/libs/*.jar /usr/src

CMD [ "java", "-jar", "/usr/src/app.jar" ]
