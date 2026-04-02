FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
COPY tools/api_doc_gen ./tools/api_doc_gen

RUN go build -o /out/api_doc_gen ./tools/api_doc_gen

FROM alpine:3.20

WORKDIR /workspace

COPY --from=build /out/api_doc_gen /usr/local/bin/api_doc_gen

ENTRYPOINT ["api_doc_gen"]
