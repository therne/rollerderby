FROM golang:1.15-alpine as base

# 0. Install Dependencies
RUN apk --no-cache --update --available upgrade
RUN apk add --no-cache make git bash

# 1. Fetch and cache go module dependencies
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# 2. Copy rest of the sources and build it
FROM base AS builder
COPY . .
RUN make

# 3. Pull binary into a clean alpine container
FROM alpine:latest
COPY --from=builder /app/build/rollerderby /usr/local/bin

ENTRYPOINT ["rollerderby"]
