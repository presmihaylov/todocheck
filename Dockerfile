FROM golang:1.26.10-alpine3.24 AS build
ARG TODOCHECK_VERSION=custom_version

ADD . /usr/src/todocheck
WORKDIR /usr/src/todocheck 

# Build
RUN go build -ldflags "-X main.version=$TODOCHECK_VERSION"

FROM alpine:3.24
COPY --from=build /usr/src/todocheck/todocheck /usr/local/bin/todocheck
ENTRYPOINT ["/usr/local/bin/todocheck"]
