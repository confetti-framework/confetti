FROM golang:1.22-alpine as project
LABEL for_development_only="false"

ENV TZ=UTC
ENV GOFLAGS="-buildvcs=false"

WORKDIR /src

RUN go install github.com/cespare/reflex@latest

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

CMD /go/bin/reflex -r '(\.go$|\.gohtml$|go\.mod$|\.env$)' -s -- sh -c "\
    go run cmd/api/main.go api:serve"
