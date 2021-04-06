FROM golang:latest
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apt-get -y update
RUN apt-get -y install git
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD ["go", "test", "-v", "./..."]
