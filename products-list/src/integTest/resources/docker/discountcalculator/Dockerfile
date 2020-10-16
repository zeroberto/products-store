FROM alpine/git AS project-env

WORKDIR /project
RUN git clone "https://github.com/zeroberto/products-store"

FROM golang:1.15.0-alpine3.12 AS build-env

WORKDIR /usr/src/app

COPY --from=project-env /project/products-store/discount-calculator ./
ADD ./application.yml ./resources

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /usr/bin/app

FROM bash

COPY --from=build-env /usr/bin/app /usr/bin/app/
COPY --from=build-env /usr/src/app/resources /usr/bin/app/resources

ENTRYPOINT /usr/bin/app/app -fconfigPath="/usr/bin/app/resources/"
