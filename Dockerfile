#build stage
FROM golang:1.20
RUN mkdir /build
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY *.go ./ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /asciiartcontainer
LABEL Name=asciiartwebdockerize Version=0.0.1
EXPOSE 8080
CMD [sudo "/asciiartcontainer"]