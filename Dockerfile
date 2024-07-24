# Build the manager binary
FROM registry.access.redhat.com/ubi9/go-toolset:1.21.11-2.1720624888@sha256:5c948cdfd0132e982426bc9d3a81eeae66871080ef274abdde1a4a8303509188 as builder

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY vendor/ vendor/
COPY pkg/ pkg/
COPY cmd/ cmd/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o multi-platform-controller cmd/controller/main.go

# Use ubi-minimal as minimal base image to package the manager binary
# Refer to https://catalog.redhat.com/software/containers/ubi8/ubi-minimal/5c359a62bed8bd75a2c3fba8 for more details
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4-1194@sha256:104cf11d890aeb7dd5728b7d7732e175a0e4018f1bb00d2faebcc8f6bf29bd52
COPY --from=builder /opt/app-root/src/multi-platform-controller /
USER 65532:65532

ENTRYPOINT ["/multi-platform-controller"]
