CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o dist/movie-sync-server-linux-mipsle
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/movie-sync-server-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/movie-sync-server-linux-arm64
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o dist/movie-sync-server-linux-mips
GOOS=windows GOARCH=amd64 go build -o dist/movie-sync-server-win-amd64.exe