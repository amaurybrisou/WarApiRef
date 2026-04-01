FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY tools/mcp_server ./tools/mcp_server

RUN go build -o /out/mcp_server ./tools/mcp_server

FROM alpine:3.20

WORKDIR /workspace

COPY --from=build /out/mcp_server /usr/local/bin/mcp_server

ENTRYPOINT ["mcp_server"]
