FROM golang:1.19.2-alpine3.15 as builder

WORKDIR /build
COPY go.mod ./
RUN go mod download && go mod tidy
COPY . .
RUN go build cmd/todofiber/main.go


# Input parameters for the Dockerfile expected in os.Getenv




FROM golang:1.19.2-alpine3.15
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"


COPY --from=builder /build/main .
COPY --from=builder /build/views ./views
COPY --from=builder /build/public ./public

EXPOSE 3000

CMD ["./main"]