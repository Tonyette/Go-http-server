#build stage
FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /go/src/Go-http-server

RUN ls -a

COPY server/ .

COPY .git .

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
  go build -ldflags "-X main.GitCommit=$GIT_COMMIT"

RUN ls

#final stage
FROM golang:alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/Go-http-server /go/src/

RUN chmod +x /go/src/Go-http-server

CMD ["/Go-http-server"]

EXPOSE 3000
