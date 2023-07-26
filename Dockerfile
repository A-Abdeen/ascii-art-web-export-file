#build stage
FROM golang:alpine AS builder
RUN mkdir /build
WORKDIR /build
RUN go get -d -v ./...
RUN cd /build && git clone 

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=asciiartwebdockerize Version=0.0.1
EXPOSE 8080
