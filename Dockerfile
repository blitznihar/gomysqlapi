FROM golang:latest

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
COPY . ./
RUN ls -la ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /gomysqlapi


EXPOSE 8081

CMD ["/gomysqlapi"]
