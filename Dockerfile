FROM golang:1.12.4

EXPOSE 8888
EXPOSE 8085

WORKDIR /goproxy

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build cmd/proxy/main.go   

CMD ./main & go run cmd/repeater/repeater.go