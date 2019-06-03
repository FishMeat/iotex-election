FROM golang:1.12.5-stretch

ENV GO111MODULE on

WORKDIR /go/src/iotex-election

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN rm -rf ./bin/server && \
    rm -rf election.db && \
    go build -o ./bin/server -v . && \
    cp ./bin/server /usr/local/bin/iotex-server  && \
    mkdir -p /etc/iotex/ && \
    cp server.yaml /etc/iotex/server.yaml && \
    rm -rf /go/src/iotex-election

CMD [ "iotex-server", "-config=/etc/iotex/server.yaml"]
