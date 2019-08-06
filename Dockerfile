FROM golang:1.12.7 AS build-env

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mutterblack ./cmd/mutterblack/mutterblack.go

# Build runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build-env /app/mutterblack .

CMD [ "./mutterblack" ]