FROM golang as builder

WORKDIR /app
COPY . .

RUN go test ./...
RUN GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

#FROM alpine
#USER root
#RUN apk add --no-cache ca-certificates
#RUN apk update && apk upgrade
#COPY --from=builder /app/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
