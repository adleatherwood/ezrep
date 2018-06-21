FROM golang:alpine as golang
WORKDIR /go/src/app
COPY ./src .
RUN go install

FROM scratch
ENV IS_DOCKER=1
COPY --from=golang /go/bin/app /app
ENTRYPOINT ["/app"]
