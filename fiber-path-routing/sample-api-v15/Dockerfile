FROM golang:1.23.7-alpine3.21 as builder
ENV CGO_ENABLED=0

COPY . /build
# Build the service binary.
WORKDIR /build
RUN go build -ldflags "-X main.build=main"


FROM alpine:3.21.1
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"
#add user and group in the alpine image
RUN addgroup -g 1000 -S vmapi && \
  adduser -u 1000 -h /cmd -G vmapi -S vmapi
#Copy from stage 0 builder only the binary files
COPY --from=builder --chown=vmapi:vmapi /build/fiber-path-routing /cmd
WORKDIR /cmd
USER vmapi
RUN ls -lR
CMD ["./fiber-path-routing"]
