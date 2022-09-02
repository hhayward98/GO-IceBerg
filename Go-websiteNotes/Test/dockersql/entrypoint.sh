wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# watches for .go file changes and recompiles main.go 
CompileDaemon --build="go build -o main main.go" --command=./main