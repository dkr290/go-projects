FROM golang:1.20-alpine3.18 AS build
RUN apk add --no-cache git
WORKDIR /tmp/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./out/holidayreservation ./cmd/web/*.go 

## very secure docker file
FROM alpine:3.18.0
RUN addgroup -S k8sgroup && adduser -S k8suser -G k8sgroup -h /home/k8suser
WORKDIr /home/k8suser
COPY --from=build /tmp/app/out/holidayreservation /home/k8suser/holidayreservation
COPY --from=build /tmp/app/templates /home/k8suser/templates
RUN chmod a-w /etc && rm -rf /bin/*
USER k8suser
CMD [ "./holidayreservation" ]