FROM golang:1.11

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
COPY . $GOPATH/src/go-api-boilerplate
WORKDIR $GOPATH/src/go-api-boilerplate
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
RUN go build -o app .

EXPOSE 8080

ENTRYPOINT ["./app"]
