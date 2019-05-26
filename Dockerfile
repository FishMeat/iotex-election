FROM golang:1.12.5-stretch AS build-env

WORKDIR apps/iotex-election

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN rm -rf ./bin/server && \
    rm -rf election.db && \
    go build -o ./bin/server -v . && \
    cp ./bin/server /usr/local/bin/iotex-server  && \
    mkdir -p /etc/iotex/ && \
    cp server.yaml /etc/iotex/server.yaml && \
    rm -rf apps/iotex-election

FROM scratch

COPY --from=build-env /etc/iotex/server.yaml /etc/iotex/server.yaml
COPY --from=build-env /usr/local/bin/iotex-server /usr/local/bin/iotex-server

ENTRYPOINT [ "/usr/local/bin/iotex-server", "-config=/etc/iotex/server.yaml"]
