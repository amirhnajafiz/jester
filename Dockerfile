# build stage
FROM golang:1.17-alpine AS builder

# set the work dir
WORKDIR /go/src/app

# copy all files
COPY . /go/src/app

# build the application
RUN go build -ldflags="-w -s" -o stan_gee

# final stage
FROM alpine

# environment variables
ENV TZ=UTC \
    PATH="/app:${PATH}"

# logs directory
RUN mkdir -p /var/log && \
    chgrp -R 0 /var/log && \
    chmod -R g=u /var/log \

# workdir to app
WORKDIR /app

# copy files from stage
COPY --from=builder /go/src/app/stan_gee /app/stan_gee

# run
ENTRYPOINT /app/stan_gee
