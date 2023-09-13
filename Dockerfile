FROM golang:1.21.1-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go

##############################
###### Production image ######
##############################

FROM alpine:3.18 AS production

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]