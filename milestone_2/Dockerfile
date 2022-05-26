FROM golang:1.16-alpine as builder
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download
COPY . /app
RUN go build -o build/app

FROM alpine
COPY --from=builder /app /usr/bin/app
EXPOSE 8080:8080
CMD [ "/usr/bin/app" ]