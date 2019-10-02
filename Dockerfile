FROM golang:1.12.4

EXPOSE 8888
EXPOSE 8085

WORKDIR /goproxy

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD go run cmd/proxy/main.go  & go run cmd/repeater/repeater.go