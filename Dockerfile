#build stage
FROM golang:alpine AS builder

RUN apk add --no-cache git

ENV ADDR=":3000"

WORKDIR /go/src/

#ENV CGO_ENABLED=0

COPY ["server/", "."]

COPY [".git", "."]

#RUN go mod download

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
  go build -o Go-http-server -ldflags "-X main.GitCommit=$GIT_COMMIT"
  
CMD ["./Go-http-server"]


