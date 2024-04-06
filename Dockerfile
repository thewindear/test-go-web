FROM golang:1.22.0-alpine3.19
WORKDIR /app
RUN apk add gcc g++ make cmake openssl-dev libtool
COPY ./ ./
RUN go build -o /app/app
EXPOSE 8080
CMD ["/app/app"]