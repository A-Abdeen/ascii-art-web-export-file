docker build -t ascii-art-web-dockerize .
docker container run -p 8080:8080 --detach --name dx ascii-art-web-dockerize