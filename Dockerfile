FROM golang:1.9-alpine as golang
ADD . /go/src/github.com/lionell/pgapps
RUN apk add --no-cache git \
    && go get github.com/lionell/pgapps/... \
    && apk del git

FROM alpine:latest
COPY --from=golang /go/bin/kubernetes app
COPY --from=golang /go/src/github.com/lionell/pgapps/res/websocket.html .
ENV PORT 8080
CMD ["./app"]
