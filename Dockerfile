ARG GO_VERSION=1.18.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src 

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
  -installsuffix 'static' \
  -o /backend-utec-inscriptions

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

COPY .env ./
COPY --from=builder /backend-utec-inscriptions /backend-utec-inscriptions

EXPOSE 3000

ENTRYPOINT [ "/backend-utec-inscriptions" ]