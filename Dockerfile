FROM golang:1.21.5-alpine AS builder
# create a build directory
RUN mkdir /build
# add all files 
ADD . /build/
WORKDIR /build/app
RUN go build

FROM alpine
RUN adduser -S -D -H -h /pingpong appuser
USER appuser
COPY --from=builder /build/app/ /pingpong/
WORKDIR /pingpong
CMD ["./app"]
