FROM golang:1.17.7-alpine3.15
LABEL maintainer="Jakkso"
WORKDIR /app
COPY go.mod ./
RUN go mod download
ADD . .
RUN go build cmd/quickshare/main.go
CMD ["./main"]