FROM  golang:1.21

ADD  .  /src

WORKDIR /src

RUN  go build ./web/server.go

ENTRYPOINT  /src/server
