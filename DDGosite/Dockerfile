# Build Stage
# First pull Golang image
FROM golang:1.18

#create dir
RUN mkdir /DDapp

# add all local files to DDapp
ADD . /DDapp

WORKDIR /DDapp

COPY go.* ./

RUN go mod download && go mod verify

#RUN go run app.go
RUN go build -o app .

EXPOSE 9000

CMD ["/DDapp/app"]

