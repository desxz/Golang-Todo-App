FROM golang:1.18-alpine
RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./ 

RUN go build -o todo-app-server

EXPOSE 5000

CMD ["./todo-app-server"]