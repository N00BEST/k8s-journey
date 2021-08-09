# Build the manager binary
ARG build_base_image=golang:latest
ARG run_base_image=alpine:latest

FROM ${build_base_image} as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY books.go books.go
COPY database.go database.go
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app

# Use distroless as minimal base image to package the app binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ${run_base_image}
WORKDIR /
COPY --from=builder /workspace/app .
USER 65532:65532

ENTRYPOINT ["/app"]