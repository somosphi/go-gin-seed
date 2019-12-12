# build stage
FROM golang:alpine as builder

RUN apk add -U make

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY Makefile .

RUN make download

COPY . .

RUN make linux BINARY=httpserver

# final stage
FROM scratch
COPY --from=builder /app/httpserver /app/
EXPOSE 3000
ENTRYPOINT ["/app/httpserver"]
