# builder image
FROM golang:1.22-bookworm as builder
RUN mkdir /build
COPY *.go /build/
COPY go.* /build/
WORKDIR /build
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .


# generate clean, final image for end users
FROM scratch
COPY --from=builder /build/server .

# executable
ENTRYPOINT [ "./server" ]
