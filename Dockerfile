# BUILD STAGE
FROM golang:1.14-alpine AS builder
WORKDIR /workspace
COPY ./ ./

RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/golang-docker ./

# FINAL STAGE
FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/golang-docker /usr/local/bin/

RUN chown -R nobody:nogroup /usr/local/bin/golang-docker
USER nobody
EXPOSE 8080
CMD [ "golang-docker" ]
