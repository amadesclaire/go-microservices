# Build code on one image 
# FROM golang:1.18-alpine as builder 
# RUN mkdir /app
# COPY . /app
# WORKDIR /app 
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
# RUN chmod +x /app/brokerApp

# New image with just the exe 
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY brokerApp /app

CMD ["/app/brokerApp"]