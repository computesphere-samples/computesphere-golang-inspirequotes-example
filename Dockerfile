# Build stage: Go 1.26.11 carries the 2026-11 Go security release (Go 1.22 is
# no longer supported). Base images are pinned by digest so a stale cached tag
# cannot bring back an unfixed build.
FROM golang:1.26.11@sha256:4165ee6bf3b7811fd7b40373290405c83d0167dcb87ee5f29c427cdcef45ff8f AS build

WORKDIR /src

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code and build a static binary
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /docker-main

# Run stage: distroless static has no shell or package manager, and the
# nonroot tag runs the app as an unprivileged user (uid 65532).
FROM gcr.io/distroless/static-debian13:nonroot@sha256:aafe993665bf6200db38a02c11a9351e782a84b6a45e7f33b6695638c88e3931

WORKDIR /app

# The app reads its templates and assets relative to /app
COPY --from=build /docker-main /docker-main
COPY templates/ ./templates/
COPY assets/ ./assets/

USER nonroot:nonroot

# The application listens on PORT, or 8080 when it isn't set
EXPOSE 8080

CMD ["/docker-main"]
