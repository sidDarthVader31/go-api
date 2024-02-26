FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o go-user-api
EXPOSE 8080
CMD ["/go-user-api"]
