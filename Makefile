preview:
	pkgsite -http localhost:8000

test:
	go test ./... -cover

clean:
	go mod tidy
	go clean -testcache
