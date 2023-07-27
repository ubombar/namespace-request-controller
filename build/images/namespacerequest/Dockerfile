FROM golang:1.20.4-alpine AS build

WORKDIR /ubombar

COPY go.mod .
RUN go mod download

COPY . ./
RUN rm -rf vendor
ENV CGO_ENABLED=0
RUN go build -o namespacerequest ./cmd/namespacerequest/

FROM alpine:3.16.2

RUN adduser -D -u 8118 ubombar
USER ubombar:ubombar

WORKDIR /ubombar/namespacerequest/
COPY --from=build --chown=ubombar:ubombar /ubombar/namespacerequest ./

CMD ["./namespacerequest"]