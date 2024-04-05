FROM golang:1.22.0-alpine3.19
WORKDIR /app
RUN apk add gcc g++ make cmake openssl-dev libtool
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
COPY ./ ./
RUN go build -o /app/app
EXPOSE 8080
CMD ["/app/app"]