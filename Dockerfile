FROM golang:1.16 AS build
WORKDIR /go/src
ENV CGO_ENABLED=0

COPY go.* ./

RUN go get -d -v ./...

COPY main.go .

COPY ./pkg ./pkg

RUN go build -a -installsuffix cgo -o notifier .

FROM scratch AS runtime
ENTRYPOINT ["./notifier"]

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/notifier ./
