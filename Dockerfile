FROM golang:1.15-alpine AS builder

ARG VERSION=0.0.1
ARG TARGET_ARCH
ENV APP_NAME=go-gena
ENV APP_VERSION=$VERSION

# GO goods
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=$TARGET_ARCH

# Move to working directory /build
WORKDIR /build 

# Copy and download dependency using go mod
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o ${APP_NAME} .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /app 

# Copy binary from build to main folder
RUN cp /build/${APP_NAME} .

# Build a small image
FROM scratch

COPY --from=builder /app/${APP_NAME} /

CMD ["/go-gena"]
