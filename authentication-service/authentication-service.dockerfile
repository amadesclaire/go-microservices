# Build code on one image 
# FROM golang:1.18-alpine as builder 
# RUN mkdir /app
# COPY . /app
# WORKDIR /app 
# RUN CGO_ENABLED=0 go build -o authApp ./cmd/api
# RUN chmod +x /app/authApp

# New image with just the exe 
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/authApp /app
COPY authApp /app

CMD ["/app/authApp"]