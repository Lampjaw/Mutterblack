FROM golang:1.12.7-alpine AS build-env

RUN apk add -U --no-cache build-base git

RUN mkdir /build
RUN mkdir /bot
WORKDIR /bot

ADD . .

RUN go get -d ./... && \
    go build -v -o /build/bot ./cmd/mutterblack

FROM alpine:latest

RUN apk add -U --no-cache iputils ca-certificates

COPY --from=build-env /build /bin

CMD [ "/bin/bot" ]