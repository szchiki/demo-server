FROM golang as builder

WORKDIR /app

ADD . ./

RUN go test -v ./... | exit 0

RUN GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

FROM alpine

EXPOSE 8080

RUN apk add --no-cache ca-certificates

RUN apk update && apk upgrade

COPY --from=builder /app/server /server

ENTRYPOINT ["/server"]
