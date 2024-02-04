FROM golang:1.20.13-alpine3.19 as builder

WORKDIR /build
COPY go.mod ./
RUN go mod download && go mod tidy
COPY . .
RUN go build cmd/app/main.go 




# Input parameters for the Dockerfile expected in os.Getenv




FROM golang:1.20.13-alpine3.19
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"
RUN addgroup -S pipeline && adduser -S  k8s-pipeline --uid 1500 -G pipeline

WORKDIR /home/k8s-pipeline
COPY --from=builder /build/main .
COPY --from=builder /build/templates ./templates
RUN ls -lR

USER k8s-pipeline

EXPOSE 8080

CMD ["./main"]