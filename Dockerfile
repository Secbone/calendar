# build container
FROM golang:latest as builder
LABEL maintainer="Secbone <secbone@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# runtime container
FROM alpine:latest  

WORKDIR /app

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main"] 