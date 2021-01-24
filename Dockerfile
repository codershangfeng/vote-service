FROM golang:1.15 as builder

WORKDIR /workspace/vote-service

COPY . .
RUN make install-swagger
RUN go mod tidy
RUN go mod verify
RUN make build


FROM alpine

LABEL Name="vote-service"

WORKDIR /root

COPY --from=builder /workspace/vote-service/bin/vote-service .

ENTRYPOINT [ "./vote-service" ]
