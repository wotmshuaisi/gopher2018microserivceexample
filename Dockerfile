FROM golang:1.11.1-alpine3.8 AS buildenv

# Allow `Golang` to retrive the dependencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 envgolang
RUN mkdir /build/ && chown envgolang /build/
USER envgolang

WORKDIR /build/
ADD . /build/

# Compile the binary, we don't want to run the cgo reslover
RUN CGO_ENABLED=0 go build -o /build/main .

# Final stage
FROM alpine:3.8


# Secure against running as root
RUN adduser -D -u 1000 runner

# Make service directory
RUN mkdir -p /service/certs
RUN mkdir /service/log

# Copy program key & cert
COPY --from=buildenv /build/main /service/main
COPY --from=buildenv /build/certs/* /service/certs/

# Change Onwer
RUN chown -R runner /service
USER runner

WORKDIR /service

EXPOSE 8080

CMD [ "/service/main" ]
