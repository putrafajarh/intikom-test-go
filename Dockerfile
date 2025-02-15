FROM golang:1.22-alpine


WORKDIR /app

# Copy only the go.mod file to install dependencies efficiently and leverage layer caching
COPY go.mod ./

# Set the GIN_MODE environment variable to release
ENV GIN_MODE=release

# Use cache mounts to speed up the installation of existing dependencies
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -o intikom-test-go

## ONLY FOR LOCAL DEVELOPMENT
COPY .env.docker .env

# Expose the port that the application will run on
EXPOSE 8080

CMD ["./intikom-test-go"]