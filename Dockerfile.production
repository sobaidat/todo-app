FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/todoapp
COPY . .
RUN go mod init todoserver
RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/todo-web-app ./todoserver.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/todoapp/bin /go/bin
EXPOSE 8000
ENTRYPOINT /go/bin/todoapp-web-app --port 8000