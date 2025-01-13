# Build the manager binary
FROM registry.access.redhat.com/ubi9/go-toolset:1.22.9-1734626445@sha256:ead35188c5748efe2b9420352aba56b02b43d8fcd7e879cc96c6b9ac2548e454 as builder

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
COPY pkg/ pkg/
COPY cmd/ cmd/

# Build
RUN GOTOOLCHAIN=auto CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o multi-platform-controller cmd/controller/main.go

# Use ubi-minimal as minimal base image to package the manager binary
# Refer to https://catalog.redhat.com/software/containers/ubi8/ubi-minimal/5c359a62bed8bd75a2c3fba8 for more details
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.5-1736404155@sha256:b87097994ed62fbf1de70bc75debe8dacf3ea6e00dd577d74503ef66452c59d6
COPY --from=builder /opt/app-root/src/multi-platform-controller /
USER 65532:65532

ENTRYPOINT ["/multi-platform-controller"]
