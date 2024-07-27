FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
COPY .env .
RUN go build -o main ./server
EXPOSE 8080
CMD [ "./main" ]