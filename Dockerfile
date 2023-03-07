FROM golang:alpine

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o blogs-api .
EXPOSE 8080
CMD ["./blogs-api"]