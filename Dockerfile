FROM golang:1.22.0-alpine3.19
WORKDIR /app
COPY ./ ./
RUN go install github.com/goreleaser/goreleaser@latest
RUN goreleaser build --single-target --snapshot
EXPOSE 8080
CMD ["/dist/binapp"]