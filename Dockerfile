ARG TARGETOS=linux
ARG TARGETARCH=amd64

ARG BUILDBASEIMAGE=golang:alpine
ARG TARGETBASEIMAGE=alpine:latest

# Building Golang image
FROM --platform=${BUILDPLATFORM:-linux/amd64} ${BUILDBASEIMAGE} AS build

# Disable CGO
ENV CGO_ENABLED=0

# Get all remaining code
RUN mkdir -p /go/src/github.com/cryptellation/codechecker
COPY ./ /go/src/github.com/cryptellation/codechecker

# Set the workdir
WORKDIR /go/src/github.com/cryptellation/codechecker

# Build everything in cmd/
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go install .

# Get final base image
FROM --platform=${TARGETOS}/${TARGETARCH} ${TARGETBASEIMAGE} AS final

# Get binary
COPY --from=build /go/bin/* /usr/local/bin