FROM golang:1.12.5-stretch AS builder

ENV GO111MODULE on

WORKDIR /go/src/iotex-election

COPY . .

RUN go mod tidy

RUN go build -o ./bin/server -v .

FROM scratch
COPY --from=builder /go/src/iotex-election/bin/server .
COPY --from=builder /go/src/iotex-election/server.yaml .

CMD [ "/server", "-config=/server.yaml"]
