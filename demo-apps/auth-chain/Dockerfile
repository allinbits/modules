FROM golang:alpine

RUN apk update
RUN apk add --no-cache ca-certificates curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /auth-chain

COPY . .

RUN go get github.com/go-delve/delve/cmd/dlv
RUN make install-debug

RUN cp ./auth-chaind /usr/bin
RUN cp ./auth-chaincli /usr/bin


EXPOSE 26656 26657 1317 9090 40000
