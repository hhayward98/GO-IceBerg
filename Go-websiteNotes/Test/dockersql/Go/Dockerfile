
FROM golang:1.15

# add all local files to DDapp
ADD . /app

WORKDIR /app

COPY go.* ./

RUN go mod download && go mod verify

#RUN go run app.go
RUN go build -o app .

EXPOSE 8080



