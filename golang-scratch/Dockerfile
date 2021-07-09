FROM golang:1.16.5-alpine3.14 as build
ADD . /golang-docker-base
WORKDIR /golang-docker-base

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch
COPY --from=build /golang-docker-base/server .

CMD ["./server"]