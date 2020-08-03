# Start by building the application.
FROM golang:1.14-buster as build

WORKDIR /go/src/foobarqix
ADD . /go/src/foobarqix

RUN go mod download

RUN go build -o /go/bin/foobarqix cmd/foobarqix/main.go

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/foobarqix /
CMD ["/foobarqix"]
