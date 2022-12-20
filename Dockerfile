# This is only for the main service, for other services please check out docker-compose.yml
FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/zenpk/dorm-system

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Set up proxy
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# Download all the dependencies
RUN go mod download

# Install the package
RUN make -C cmd build

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["bin/main", "-mode", "prod"]
