#docker file for deployment
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go.mod download

COPY . .

RUN go build -0 main .

EXPOSE 8080

CMD [ "./main" ]