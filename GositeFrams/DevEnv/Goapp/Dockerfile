FROM golang:1.19.4

WORKDIR /go/src/app
COPY . .


RUN go mod download && go mod verify



#Used for recompiling program after file edits
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

EXPOSE 8088