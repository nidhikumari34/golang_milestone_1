FROM golang:latest
LABEL maintainer="Quique <hello@pragmaticreviews.com>"
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 8000
RUN go build
CMD ["./show"]

