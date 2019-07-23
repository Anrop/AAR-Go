FROM golang:1.12-alpine AS builder
RUN apk --no-cache add git
WORKDIR /build/
COPY . /build/
RUN GO111MODULE=on go build -mod=vendor

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /build/AAR-Go /app/aar
CMD /app/aar
