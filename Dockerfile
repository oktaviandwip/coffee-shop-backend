FROM golang:latest AS build

WORKDIR /goapp

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /goapp/goback ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /goapp /app

ENV PATH="/app:${PATH}"

EXPOSE 8081

ENTRYPOINT [ "goback" ]