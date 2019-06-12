FROM golang:latest

COPY ./ /go/src/github.com/elastic/integrations-registry
EXPOSE 8080

WORKDIR "/go/src/github.com/elastic/integrations-registry"
RUN GO111MODULE=on go mod vendor
<<<<<<< HEAD

# Make sure it's accessible from outside the container
ENTRYPOINT ["go", "run", "main.go", "--address=0.0.0.0:8080"]
