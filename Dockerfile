FROM golang:1.22.2-alpine as builder

WORKDIR /workspace

# add go modules lockfiles
COPY go.mod go.sum ./
RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
RUN go run github.com/steebchen/prisma-client-go prefetch

COPY ./ ./
# generate the Prisma Client Go client
RUN go generate ./...

# build a fully standalone binary with zero dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .


# <!> railway.app does not really have a way to run a migration container on the volume, so we have to improvise here
# in order to run migrations, we need to have go and prisma installed
FROM debian:bookworm-slim

# set timezone and copy certs
ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy app binary
COPY --from=builder /workspace/app /app

# copy prisma binaries
COPY --from=builder /root/.cache/prisma/binaries /root/.cache/prisma/binaries