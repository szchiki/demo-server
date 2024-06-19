FROM golang as builder

WORKDIR /app
COPY . .

RUN go test ./...
RUN GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

#FROM alpine # Makes our image a lot lighter, cruicial for mass deployment
#USER root # Don't use root, but the user may need to be explicitly set for the alpine image.
#RUN apk add --no-cache ca-certificates
#RUN apk update && apk upgrade
#COPY --from=builder /app/server /server

EXPOSE 8080
ENTRYPOINT ["/app/server"]
