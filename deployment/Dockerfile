FROM golang:alpine

RUN apk update && apk add --no-cache git make

WORKDIR /app

COPY . /app

RUN make clean && make install

RUN make build

ENTRYPOINT ["/app/binary serve"]