ARG GO_VERSION=1.12.9
FROM golang:${GO_VERSION}-alpine AS builder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git socat
WORKDIR /src
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app .

FROM scratch AS final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app
EXPOSE 8080
USER nobody:nobody
ENTRYPOINT ["/app"]
