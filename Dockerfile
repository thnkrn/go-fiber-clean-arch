FROM golang:1.20.4-alpine3.16 AS build_base
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-fiber-clean-arch

# Populate the module cache based on the go.{mod,sum} files.
COPY . .
RUN go get -u -t -d -v ./... && go mod download && go mod tidy && go mod vendor

# Build the Go app
RUN go build -o ./out/go-fiber-clean-arch -v ./cmd/api

# Start fresh from a smaller image
FROM alpine:3.16
RUN apk add --no-cache ca-certificates iwatch=0.2.2-r0
COPY --from=build_base /tmp/go-fiber-clean-arch/out/go-fiber-clean-arch /app/go-fiber-clean-arch
COPY ./.env /app

USER root
# This container exposes port 8080 to the outside world
EXPOSE 8080

WORKDIR /app

# Init fake process to grab PID 1 before the Fiber app does, https://github.com/gofiber/fiber/issues/1036
RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Run the binary program produced by `go install`
CMD ["./go-fiber-clean-arch"]
