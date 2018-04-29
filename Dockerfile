FROM golang:1.8.3-alpine

WORKDIR /app

RUN apk add --no-cache bash git openssh

EXPOSE 8080

# Now just add the binary
ADD . /app/

RUN go get github.com/gorilla/mux
RUN go get github.com/Silvian/GoRestApi/api
RUN go get github.com/Silvian/GoRestApi/db

RUN go build -o /app/main cmd/main.go

CMD ["go", "run", "cmd/main.go"]
