FROM golang:1.19
LABEL authors="ahmed, ahmed"
LABEL version="0.1"
COPY . /ascii-art-web-dockerize
WORKDIR /ascii-art-web-dockerize/webprogram
RUN go build -v main.go
CMD ./main
EXPOSE 8080