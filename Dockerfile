FROM golang:1.21-alpine AS buildStage
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
# RUN CGO_ENABLED=0 go test -v --tags=unit ./...
RUN go build -o ./out .

# ===================

FROM alpine:3.16.2
COPY --from=buildStage /app/out /app
CMD ["/app"]