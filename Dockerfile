FROM golang:1.25-alpine3.22 AS build-stage

RUN apk add --no-cache git 

WORKDIR /app

ENV CGO_ENABLED=0
ARG VERSION=dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w -X main.VERSION=${VERSION}" \
    -o pg-mcp cmd/pg-mcp/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=build-stage /app/pg-mcp .

ENTRYPOINT [ "./pg-mcp" ]
CMD [ "-version" ]
