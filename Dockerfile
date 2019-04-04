# https://medium.com/@petomalina/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888
FROM golang:1.12 as build-env
ENV GO111MODULE=on
WORKDIR /app
COPY . .
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/server ./cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/client ./cmd/client/main.go
FROM alpine AS last-phase
RUN apk add ca-certificates
# Finally we copy the statically compiled Go binary.
COPY --from=build-env /go/bin/server /bin/server
COPY --from=build-env /go/bin/client /bin/client
CMD ["/bin/server"]