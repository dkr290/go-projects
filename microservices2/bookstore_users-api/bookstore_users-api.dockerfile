FROM golang:1.20.2-alpine3.17 as builder


WORKDIR /build
COPY go.mod ./
RUN go mod download && go mod tidy
COPY . .

RUN CGO_ENABLED=0 go build -o bookstore-usersapi main.go
RUN ls -l



FROM golang:1.20.2-alpine3.17

RUN chmod a-w /etc

RUN addgroup -S pipeline && adduser -S  k8s-pipeline --uid 1500 -G pipeline -h /home/k8s-pipeline

WORKDIR /home/k8s-pipeline
COPY --from=builder /build/bookstore-usersapi .
COPY .env . 

RUN ls -l

USER k8s-pipeline

EXPOSE 8080

CMD ["/home/k8s-pipeline/bookstore-usersapi"]