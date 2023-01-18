FROM golang:1.18.3 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o server ./services/buildingdata-access

FROM scratch as server
WORKDIR /app

FROM mongo:latest as mongo

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server .

ENTRYPOINT ["./server", "9002"]
