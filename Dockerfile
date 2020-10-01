FROM golang:latest as builder
WORKDIR /app/backend
ADD . /app/backend
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /gateway_srv /app/backend/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /gateway_srv ./
RUN chmod +x ./gateway_srv
ENTRYPOINT ["./gateway_srv"]


