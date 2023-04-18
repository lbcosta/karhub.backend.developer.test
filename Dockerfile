FROM golang:1.20

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

CMD [ "/app/server" ]

