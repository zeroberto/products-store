FROM golang:1.15.0-alpine3.12 AS build-env

ADD . /usr/src/app
WORKDIR /usr/src/app

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /usr/bin/app

FROM scratch

COPY --from=build-env /usr/bin/app ./

ENTRYPOINT [ "/app" ]
