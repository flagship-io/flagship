FROM golang:1.18-alpine as builder
RUN apk add --update make
WORKDIR /go/src/github/flagship-io/flagship

ARG FLAGSHIP_VERSION
ENV FLAGSHIP_VERSION $FLAGSHIP_VERSION

# Download dependencies before building
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github/flagship-io/flagship/flagship ./
CMD ["/bin/sh"]