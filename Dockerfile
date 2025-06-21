FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main .

EXPOSE 3000

ENTRYPOINT ["/app/main"]


