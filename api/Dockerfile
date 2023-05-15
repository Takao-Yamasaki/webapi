FROM golang:1.20

WORKDIR /api
COPY main.go go.mod go.sum /api/
RUN go build main.go

CMD ["./main"]
