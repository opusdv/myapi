FROM golang:latest as builder
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 8080
RUN go build -v -o app cmd/main.go

FROM debian:buster-slim
COPY --from=builder /app/app /usr/bin/app
ENTRYPOINT [ "/usr/bin/app" ]