##
##ENVS
#KUBERNETES_NAMEPSACE
#KUBERNETES_MAXTIME
#GOCRON_SCHEDULE

FROM golang:1.19.2-alpine3.15 as builder

WORKDIR /build
COPY go.mod ./
RUN go mod download && go mod tidy
COPY . .
RUN go build cmd/api/*.go -o main


# Input parameters for the Dockerfile expected in os.Getenv




FROM golang:1.19.2-alpine3.15
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"
RUN addgroup -S pipeline && adduser -S  k8s-pipeline --uid 1500 -G pipeline

WORKDIR /home/k8s-pipeline
COPY --from=builder /build/main .

USER k8s-pipeline

CMD ["./main"]