FROM golang:latest

WORKDIR /cmd/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

#Port Environment variable to determine port for listening
ENV PORT=8080

EXPOSE 8080

CMD ["go", "run", "cmd/app/main.go"]
