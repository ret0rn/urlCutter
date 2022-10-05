FROM golang:alpine AS builder

WORKDIR /build/
COPY . .
RUN go mod download
RUN go build -o urlcutter main.go

FROM alpine:latest 
WORKDIR /app/
COPY --from=builder /build/urlcutter .
COPY --from=builder /build/configs/* ./configs/
EXPOSE 80
CMD ["./urlcutter"]