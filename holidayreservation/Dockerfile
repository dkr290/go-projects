FROM golang:1.20-alpine3.18 AS build
RUN apk add --no-cache git
WORKDIR /tmp/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./out/holidayreservation ./cmd/web/*.go 


FROM alpine:3.18.0
RUN apk add ca-certificates
WORKDIR /app
COPY --from=build /tmp/app/out/holidayreservation /app/holidayreservation
COPY --from=build /tmp/app/templates /app/templates
RUN ls -lR /app
CMD [ "./holidayreservation" ]