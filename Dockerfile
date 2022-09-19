FROM golang:1.19 as build

WORKDIR /build
COPY demo demo
COPY go.mod go.sum ./

RUN CGO_ENABLED=0 go build -o built_app_sender demo/cmd/sender/main.go && \
    CGO_ENABLED=0 go build -o built_app_receiver demo/cmd/receiver/main.go

FROM alpine:3.16 as app_sender
COPY --from=build /build/built_app_sender .
ENTRYPOINT ["./built_app_sender"]

FROM alpine:3.16 as app_receiver
COPY --from=build /build/built_app_receiver .
ENTRYPOINT ["./built_app_receiver"]
