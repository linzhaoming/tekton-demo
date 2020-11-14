FROM golang:1.15 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/linzhaoming/tekton-demo
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o startkit

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/linzhaoming/tekton-demo/demo /demo

# Run the web service on container startup.
CMD ["/demo"]